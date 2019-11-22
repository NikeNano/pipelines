from kfp.dsl import ContainerOp
from kfp.utils import use_secret
import unittest
import inspect


class TestAddSecrets(unittest.TestCase):
    def test_default_use_secret(self):
        spec = inspect.getfullargspec(use_secret)
        assert len(spec.defaults) == 5
        for spec in spec.defaults:
            assert spec == None


    def test_use_default_use_secret(self):
        op1 = ContainerOp(name='op1', image='image')
        secret_name = "my-secret"
        secret_path = "/here/are/my/secret"
        op1 = op1.apply(use_secret(secret_name=secret_name, 
        secret_volume_mount_path=secret_path))
        assert type(op1.container.env) == type(None)

        container_dict = op1.container.to_dict()
        volume_mounts = container_dict["volume_mounts"][0]
        assert type(volume_mounts)==dict
        assert volume_mounts["name"] == secret_name + '_volume'
        assert volume_mounts["mount_path"] == secret_path


    def test_use_set_volume_use_secret(self):
        op1 = ContainerOp(name='op1', image='image')
        secret_name = "my-secret"
        secret_path = "/here/are/my/secret"
        volume_name = "my_volume"
        op1 = op1.apply(use_secret(secret_name=secret_name, 
            secret_volume_mount_path=secret_path, 
            volume_name = volume_name))
        assert type(op1.container.env) == type(None)

        container_dict = op1.container.to_dict()
        volume_mounts = container_dict["volume_mounts"][0]
        assert type(volume_mounts)==dict
        assert volume_mounts["name"] == volume_name
        assert volume_mounts["mount_path"] == secret_path


    def test_use_set_env_ues_secret(self):
        op1 = ContainerOp(name='op1', image='image')
        secret_name = "my-secret"
        secret_path = "/here/are/my/secret/"
        env_variable = "MY_SECRET"
        secret_file_path_in_volume = "secret.json"
        op1 = op1.apply(use_secret(secret_name=secret_name, 
            secret_volume_mount_path=secret_path,
            env_variable=env_variable,
            secret_file_path_in_volume=secret_file_path_in_volume))
        assert len(op1.container.env) ==1

        container_dict = op1.container.to_dict()
        volume_mounts = container_dict["volume_mounts"][0]
        assert type(volume_mounts)==dict
        assert volume_mounts["name"] == secret_name + '_volume'
        assert volume_mounts["mount_path"] == secret_path 
        env_dict = op1.container.env[0].to_dict()
        assert env_dict["name"] == env_variable
        assert env_dict["value"] == secret_path + secret_file_path_in_volume
