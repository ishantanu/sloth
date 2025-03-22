// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apislothv1 "github.com/slok/sloth/pkg/kubernetes/api/sloth/v1"
	versioned "github.com/slok/sloth/pkg/kubernetes/gen/clientset/versioned"
	internalinterfaces "github.com/slok/sloth/pkg/kubernetes/gen/informers/externalversions/internalinterfaces"
	slothv1 "github.com/slok/sloth/pkg/kubernetes/gen/listers/sloth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PrometheusServiceLevelInformer provides access to a shared informer and lister for
// PrometheusServiceLevels.
type PrometheusServiceLevelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() slothv1.PrometheusServiceLevelLister
}

type prometheusServiceLevelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPrometheusServiceLevelInformer constructs a new informer for PrometheusServiceLevel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPrometheusServiceLevelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPrometheusServiceLevelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPrometheusServiceLevelInformer constructs a new informer for PrometheusServiceLevel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPrometheusServiceLevelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SlothV1().PrometheusServiceLevels(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SlothV1().PrometheusServiceLevels(namespace).Watch(context.TODO(), options)
			},
		},
		&apislothv1.PrometheusServiceLevel{},
		resyncPeriod,
		indexers,
	)
}

func (f *prometheusServiceLevelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPrometheusServiceLevelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *prometheusServiceLevelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apislothv1.PrometheusServiceLevel{}, f.defaultInformer)
}

func (f *prometheusServiceLevelInformer) Lister() slothv1.PrometheusServiceLevelLister {
	return slothv1.NewPrometheusServiceLevelLister(f.Informer().GetIndexer())
}
