// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterPlotWidgetDefinitionRequests Widget definition.
type ScatterPlotWidgetDefinitionRequests struct {
	// Scatterplot request containing formulas and functions.
	Table *ScatterplotTableRequest `json:"table,omitempty"`
	// Updated scatter plot.
	X *ScatterPlotRequest `json:"x,omitempty"`
	// Updated scatter plot.
	Y *ScatterPlotRequest `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewScatterPlotWidgetDefinitionRequests instantiates a new ScatterPlotWidgetDefinitionRequests object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterPlotWidgetDefinitionRequests() *ScatterPlotWidgetDefinitionRequests {
	this := ScatterPlotWidgetDefinitionRequests{}
	return &this
}

// NewScatterPlotWidgetDefinitionRequestsWithDefaults instantiates a new ScatterPlotWidgetDefinitionRequests object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterPlotWidgetDefinitionRequestsWithDefaults() *ScatterPlotWidgetDefinitionRequests {
	this := ScatterPlotWidgetDefinitionRequests{}
	return &this
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ScatterPlotWidgetDefinitionRequests) GetTable() ScatterplotTableRequest {
	if o == nil || o.Table == nil {
		var ret ScatterplotTableRequest
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterPlotWidgetDefinitionRequests) GetTableOk() (*ScatterplotTableRequest, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ScatterPlotWidgetDefinitionRequests) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given ScatterplotTableRequest and assigns it to the Table field.
func (o *ScatterPlotWidgetDefinitionRequests) SetTable(v ScatterplotTableRequest) {
	o.Table = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *ScatterPlotWidgetDefinitionRequests) GetX() ScatterPlotRequest {
	if o == nil || o.X == nil {
		var ret ScatterPlotRequest
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterPlotWidgetDefinitionRequests) GetXOk() (*ScatterPlotRequest, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *ScatterPlotWidgetDefinitionRequests) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given ScatterPlotRequest and assigns it to the X field.
func (o *ScatterPlotWidgetDefinitionRequests) SetX(v ScatterPlotRequest) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *ScatterPlotWidgetDefinitionRequests) GetY() ScatterPlotRequest {
	if o == nil || o.Y == nil {
		var ret ScatterPlotRequest
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterPlotWidgetDefinitionRequests) GetYOk() (*ScatterPlotRequest, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *ScatterPlotWidgetDefinitionRequests) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given ScatterPlotRequest and assigns it to the Y field.
func (o *ScatterPlotWidgetDefinitionRequests) SetY(v ScatterPlotRequest) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterPlotWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterPlotWidgetDefinitionRequests) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Table *ScatterplotTableRequest `json:"table,omitempty"`
		X     *ScatterPlotRequest      `json:"x,omitempty"`
		Y     *ScatterPlotRequest      `json:"y,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"table", "x", "y"})
	} else {
		return err
	}
	if all.Table != nil && all.Table.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Table = all.Table
	if all.X != nil && all.X.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.X = all.X
	if all.Y != nil && all.Y.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Y = all.Y
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
