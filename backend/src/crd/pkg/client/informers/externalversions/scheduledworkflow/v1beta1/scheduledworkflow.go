// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	scheduledworkflowv1beta1 "github.com/kubeflow/pipelines/backend/src/crd/pkg/apis/scheduledworkflow/v1beta1"
	versioned "github.com/kubeflow/pipelines/backend/src/crd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/pipelines/backend/src/crd/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubeflow/pipelines/backend/src/crd/pkg/client/listers/scheduledworkflow/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ScheduledWorkflowInformer provides access to a shared informer and lister for
// ScheduledWorkflows.
type ScheduledWorkflowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ScheduledWorkflowLister
}

type scheduledWorkflowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScheduledWorkflowInformer constructs a new informer for ScheduledWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScheduledWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScheduledWorkflowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScheduledWorkflowInformer constructs a new informer for ScheduledWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScheduledWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	ctx := context.Background()
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScheduledworkflowV1beta1().ScheduledWorkflows(namespace).List(ctx, options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScheduledworkflowV1beta1().ScheduledWorkflows(namespace).Watch(ctx, options)
			},
		},
		&scheduledworkflowv1beta1.ScheduledWorkflow{},
		resyncPeriod,
		indexers,
	)
}

func (f *scheduledWorkflowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScheduledWorkflowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scheduledWorkflowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&scheduledworkflowv1beta1.ScheduledWorkflow{}, f.defaultInformer)
}

func (f *scheduledWorkflowInformer) Lister() v1beta1.ScheduledWorkflowLister {
	return v1beta1.NewScheduledWorkflowLister(f.Informer().GetIndexer())
}
