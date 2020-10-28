package client

import (
	"context"
	"errors"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type FakePodClient struct {
}

// FakeBadPodClient
func (FakePodClient) GetEphemeralContainers(ctx context.Context, podName string, options v1.GetOptions) (*corev1.EphemeralContainers, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) UpdateEphemeralContainers(ctx context.Context, podName string, options *corev1.EphemeralContainers, opts v1.UpdateOptions) (*corev1.EphemeralContainers, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) Create(ctx context.Context, pod *corev1.Pod, opts v1.CreateOptions) (*corev1.Pod, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) Update(ctx context.Context, pod *corev1.Pod, opts v1.UpdateOptions) (*corev1.Pod, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) UpdateStatus(ctx context.Context, pod *corev1.Pod, opts v1.UpdateOptions) (*corev1.Pod, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return nil
}

// FakeBadPodClient
func (FakePodClient) DeleteCollection(ctx context.Context, options v1.DeleteOptions, listOptions v1.ListOptions) error {
	glog.Error("This fake method is not yet implemented.")
	return nil
}

// FakeBadPodClient
func (FakePodClient) Get(ctx context.Context, name string, options v1.GetOptions) (*corev1.Pod, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) List(ctx context.Context, opts v1.ListOptions) (*corev1.PodList, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Pod, err error) {
	glog.Error("This fake method is not yet implemented.")
	return nil, nil
}

// FakeBadPodClient
func (FakePodClient) Bind(ctx context.Context, binding *corev1.Binding, opts v1.CreateOptions) error {
	glog.Error("This fake method is not yet implemented.")
	return nil
}

// FakeBadPodClient
func (FakePodClient) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	glog.Error("This fake method is not yet implemented.")
	return nil
}

// FakeBadPodClient
func (FakePodClient) GetLogs(name string, opts *corev1.PodLogOptions) *rest.Request {
	glog.Error("This fake method is not yet implemented.")
	return nil
}

// FakeBadPodClient
type FakeBadPodClient struct {
	FakePodClient
}

// FakeBadPodClient
func (FakeBadPodClient) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return errors.New("failed to delete pod")
}
