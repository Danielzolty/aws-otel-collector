// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesResponse A message containing the response to a timeseries query.
type TimeseriesResponse struct {
	// The object describing a timeseries response.
	Attributes *TimeseriesResponseAttributes `json:"attributes,omitempty"`
	// The type of the resource. The value should always be timeseries_response.
	Type *TimeseriesFormulaResponseType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTimeseriesResponse instantiates a new TimeseriesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesResponse() *TimeseriesResponse {
	this := TimeseriesResponse{}
	var typeVar TimeseriesFormulaResponseType = TIMESERIESFORMULARESPONSETYPE_TIMESERIES_RESPONSE
	this.Type = &typeVar
	return &this
}

// NewTimeseriesResponseWithDefaults instantiates a new TimeseriesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesResponseWithDefaults() *TimeseriesResponse {
	this := TimeseriesResponse{}
	var typeVar TimeseriesFormulaResponseType = TIMESERIESFORMULARESPONSETYPE_TIMESERIES_RESPONSE
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TimeseriesResponse) GetAttributes() TimeseriesResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TimeseriesResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResponse) GetAttributesOk() (*TimeseriesResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TimeseriesResponse) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TimeseriesResponseAttributes and assigns it to the Attributes field.
func (o *TimeseriesResponse) SetAttributes(v TimeseriesResponseAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TimeseriesResponse) GetType() TimeseriesFormulaResponseType {
	if o == nil || o.Type == nil {
		var ret TimeseriesFormulaResponseType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResponse) GetTypeOk() (*TimeseriesFormulaResponseType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TimeseriesResponse) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given TimeseriesFormulaResponseType and assigns it to the Type field.
func (o *TimeseriesResponse) SetType(v TimeseriesFormulaResponseType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *TimeseriesResponseAttributes  `json:"attributes,omitempty"`
		Type       *TimeseriesFormulaResponseType `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Type = all.Type
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
