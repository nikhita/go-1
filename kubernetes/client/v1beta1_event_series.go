/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type V1beta1EventSeries struct {

	// Number of occurrences in this series up to the last heartbeat time
	Count int32 `json:"count"`

	// Time when last Event from the series was seen before last heartbeat.
	LastObservedTime time.Time `json:"lastObservedTime"`

	// Information whether this series is ongoing or finished.
	State string `json:"state"`
}
