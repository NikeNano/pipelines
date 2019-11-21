def use_secret(secret_name=None, volume_name=None, secret_volume_mount_path=None, env_variable=None, secret_file_path_in_volume=None):
    """    
       An operator that configures the container to use a secret.
       
       This assumes that the secret is created and availabel in the k8s cluster.
    
    Keyword Arguments:
        secret_name {String} -- The k8s secret name (default: {None})
        volume_name {String} -- The pod volume name (default: {None})
        secret_volume_mount_path {String} -- The path to the secret that is mounted  (default: {None})
        env_variable {String} -- Env variable pointing to the mounted secret file. Requires both the env_variable and secret_file_path_in_volume to be defined. 
                                 The value is the path to the secret (default: {None})
        secret_file_path_in_volume {String} -- The path to the secret in the volume. This will be the value of env_variable. 
                                 Both env_variable and secret_file_path_in_volume needs to be set if any env variable should be created (default: {None})
    
    Raises:
        ValueError: If not the necessary variables (secret_name, volume_name", secret_volume_mount_path) are supplied.
                    Or only one of  env_variable and secret_file_path_in_volume are supplied
    
    Returns:
        [ContainerOperator] -- Returns the container operator after it has been modified. 
    """

    if not volume_name:
        volumen_name = secret_name + "_volume"

    params = [secret_name, volume_name, secret_volume_mount_path]
    param_names = ["secret_name", "volume_name", "secret_volume_mount_path"]
    for param, param_name in zip(params, param_names):
        if param is None:
            raise ValueError("'{}' needs to be specified, is: {}".format(param_name, param))
    
    if (env_variable and not secret_file_path_in_volume) or (secret_file_path_in_volume and not env_variable):
        raise ValueError("Both {} and {} needs to be supplied together or not at all".format(env_variable, secret_file_path_in_volume))

    def _use_secret(task):
        import os 
        from kubernetes import client as k8s_client
        task = task.add_volume(
            k8s_client.V1Volume(
                name=volume_name,
                secret=k8s_client.V1SecretVolumeSource(
                    secret_name=secret_name
                )
            )
        ).add_volume_mount(
                k8s_client.V1VolumeMount(
                    name=volume_name,
                    mount_path=secret_volume_mount_path
                )
            )
        if env_variable:
            task.container.add_env_variable(
                k8s_client.V1EnvVar(
                    name=env_variable,
                    value=os.path.join(secret_volume_mount_path, secret_file_path_in_volume),
                )
            )
        return task
    
    return _use_secret