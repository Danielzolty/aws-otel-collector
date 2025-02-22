// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PagerDutyServiceName PagerDuty service object name.
type PagerDutyServiceName struct {
	// Your service name associated service key in PagerDuty.
	ServiceName string `json:"service_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewPagerDutyServiceName instantiates a new PagerDutyServiceName object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPagerDutyServiceName(serviceName string) *PagerDutyServiceName {
	this := PagerDutyServiceName{}
	this.ServiceName = serviceName
	return &this
}

// NewPagerDutyServiceNameWithDefaults instantiates a new PagerDutyServiceName object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPagerDutyServiceNameWithDefaults() *PagerDutyServiceName {
	this := PagerDutyServiceName{}
	return &this
}

// GetServiceName returns the ServiceName field value.
func (o *PagerDutyServiceName) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *PagerDutyServiceName) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *PagerDutyServiceName) SetServiceName(v string) {
	o.ServiceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PagerDutyServiceName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["service_name"] = o.ServiceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PagerDutyServiceName) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ServiceName *string `json:"service_name"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field service_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"service_name"})
	} else {
		return err
	}
	o.ServiceName = *all.ServiceName
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
