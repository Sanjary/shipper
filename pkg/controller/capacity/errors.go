package capacity

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
)

type MultipleOwnerReferencesError error

func NewMultipleOwnerReferencesError(name string, references int) MultipleOwnerReferencesError {
	return MultipleOwnerReferencesError(fmt.Errorf(
		"expected exactly one owner for CapacityTarget %q, got %d",
		name, references))
}

type ReleaseIsGoneError error

func NewReleaseIsGoneError(name string, expectedUID types.UID, gotUID types.UID) ReleaseIsGoneError {
	return ReleaseIsGoneError(fmt.Errorf(
		"the owner Release for CapacityTarget %q is gone; expected UID %s but got %s",
		name, expectedUID, gotUID,
	))
}

type InvalidCapacityTargetError error

func NewInvalidCapacityTargetError(releaseName string, count int) InvalidCapacityTargetError {
	var message error
	if count < 1 {
		message = fmt.Errorf("missing capacity target with release label %q", releaseName)
	} else if count > 1 {
		message = fmt.Errorf(
			"expected one capacity target for release label %q, got %d instead", releaseName, count)
	} else {
		// Since we should have only 1 Capacity Target object per release,
		// having a count of 1 here is a programmer error.
		panic("programmer error: NewInvalidCapacityTargetError() should not be called with count of 1")
	}
	return InvalidCapacityTargetError(message)
}
