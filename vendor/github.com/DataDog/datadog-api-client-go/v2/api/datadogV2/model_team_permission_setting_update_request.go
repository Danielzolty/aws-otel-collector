// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamPermissionSettingUpdateRequest Team permission setting update request
type TeamPermissionSettingUpdateRequest struct {
	// Team permission setting update
	Data TeamPermissionSettingUpdate `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTeamPermissionSettingUpdateRequest instantiates a new TeamPermissionSettingUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamPermissionSettingUpdateRequest(data TeamPermissionSettingUpdate) *TeamPermissionSettingUpdateRequest {
	this := TeamPermissionSettingUpdateRequest{}
	this.Data = data
	return &this
}

// NewTeamPermissionSettingUpdateRequestWithDefaults instantiates a new TeamPermissionSettingUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamPermissionSettingUpdateRequestWithDefaults() *TeamPermissionSettingUpdateRequest {
	this := TeamPermissionSettingUpdateRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *TeamPermissionSettingUpdateRequest) GetData() TeamPermissionSettingUpdate {
	if o == nil {
		var ret TeamPermissionSettingUpdate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingUpdateRequest) GetDataOk() (*TeamPermissionSettingUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *TeamPermissionSettingUpdateRequest) SetData(v TeamPermissionSettingUpdate) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamPermissionSettingUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamPermissionSettingUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *TeamPermissionSettingUpdate `json:"data"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = *all.Data
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
