// Package clusterclientstore provides a thread-safe storage for Kubernetes API
// clients connected to target clusters. The internal storage is updated
// automatically by watching for Cluster and Secret resources. New Cluster
// objects trigger a client/informer creation, updates to Secret objects trigger
// re-creation of a client/informer, and Cluster deletions cause the removal of
// a client and its associated informer.
//
// Important Note
//
// before using the methods for retrieving cluster-specific objects from the store, you *must* call the `Run()` method.
package clusterclientstore
