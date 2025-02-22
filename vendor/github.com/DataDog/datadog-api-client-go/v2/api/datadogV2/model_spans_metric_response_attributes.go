// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansMetricResponseAttributes The object describing a Datadog span-based metric.
type SpansMetricResponseAttributes struct {
	// The compute rule to compute the span-based metric.
	Compute *SpansMetricResponseCompute `json:"compute,omitempty"`
	// The span-based metric filter. Spans matching this filter will be aggregated in this metric.
	Filter *SpansMetricResponseFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []SpansMetricResponseGroupBy `json:"group_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSpansMetricResponseAttributes instantiates a new SpansMetricResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansMetricResponseAttributes() *SpansMetricResponseAttributes {
	this := SpansMetricResponseAttributes{}
	return &this
}

// NewSpansMetricResponseAttributesWithDefaults instantiates a new SpansMetricResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansMetricResponseAttributesWithDefaults() *SpansMetricResponseAttributes {
	this := SpansMetricResponseAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SpansMetricResponseAttributes) GetCompute() SpansMetricResponseCompute {
	if o == nil || o.Compute == nil {
		var ret SpansMetricResponseCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricResponseAttributes) GetComputeOk() (*SpansMetricResponseCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SpansMetricResponseAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given SpansMetricResponseCompute and assigns it to the Compute field.
func (o *SpansMetricResponseAttributes) SetCompute(v SpansMetricResponseCompute) {
	o.Compute = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SpansMetricResponseAttributes) GetFilter() SpansMetricResponseFilter {
	if o == nil || o.Filter == nil {
		var ret SpansMetricResponseFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricResponseAttributes) GetFilterOk() (*SpansMetricResponseFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SpansMetricResponseAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given SpansMetricResponseFilter and assigns it to the Filter field.
func (o *SpansMetricResponseAttributes) SetFilter(v SpansMetricResponseFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *SpansMetricResponseAttributes) GetGroupBy() []SpansMetricResponseGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []SpansMetricResponseGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricResponseAttributes) GetGroupByOk() (*[]SpansMetricResponseGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *SpansMetricResponseAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []SpansMetricResponseGroupBy and assigns it to the GroupBy field.
func (o *SpansMetricResponseAttributes) SetGroupBy(v []SpansMetricResponseGroupBy) {
	o.GroupBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansMetricResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansMetricResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Compute *SpansMetricResponseCompute  `json:"compute,omitempty"`
		Filter  *SpansMetricResponseFilter   `json:"filter,omitempty"`
		GroupBy []SpansMetricResponseGroupBy `json:"group_by,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "filter", "group_by"})
	} else {
		return err
	}
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Compute = all.Compute
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
