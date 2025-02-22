// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOWidgetDefinition Use the SLO and uptime widget to track your SLOs (Service Level Objectives) and uptime on screenboards and timeboards.
type SLOWidgetDefinition struct {
	// Additional filters applied to the SLO query.
	AdditionalQueryFilters *string `json:"additional_query_filters,omitempty"`
	// Defined global time target.
	GlobalTimeTarget *string `json:"global_time_target,omitempty"`
	// Defined error budget.
	ShowErrorBudget *bool `json:"show_error_budget,omitempty"`
	// ID of the SLO displayed.
	SloId *string `json:"slo_id,omitempty"`
	// Times being monitored.
	TimeWindows []WidgetTimeWindows `json:"time_windows,omitempty"`
	// Title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the SLO widget.
	Type SLOWidgetDefinitionType `json:"type"`
	// Define how you want the SLO to be displayed.
	ViewMode *WidgetViewMode `json:"view_mode,omitempty"`
	// Type of view displayed by the widget.
	ViewType string `json:"view_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSLOWidgetDefinition instantiates a new SLOWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOWidgetDefinition(typeVar SLOWidgetDefinitionType, viewType string) *SLOWidgetDefinition {
	this := SLOWidgetDefinition{}
	this.Type = typeVar
	this.ViewType = viewType
	return &this
}

// NewSLOWidgetDefinitionWithDefaults instantiates a new SLOWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOWidgetDefinitionWithDefaults() *SLOWidgetDefinition {
	this := SLOWidgetDefinition{}
	var typeVar SLOWidgetDefinitionType = SLOWIDGETDEFINITIONTYPE_SLO
	this.Type = typeVar
	var viewType string = "detail"
	this.ViewType = viewType
	return &this
}

// GetAdditionalQueryFilters returns the AdditionalQueryFilters field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetAdditionalQueryFilters() string {
	if o == nil || o.AdditionalQueryFilters == nil {
		var ret string
		return ret
	}
	return *o.AdditionalQueryFilters
}

// GetAdditionalQueryFiltersOk returns a tuple with the AdditionalQueryFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetAdditionalQueryFiltersOk() (*string, bool) {
	if o == nil || o.AdditionalQueryFilters == nil {
		return nil, false
	}
	return o.AdditionalQueryFilters, true
}

// HasAdditionalQueryFilters returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasAdditionalQueryFilters() bool {
	return o != nil && o.AdditionalQueryFilters != nil
}

// SetAdditionalQueryFilters gets a reference to the given string and assigns it to the AdditionalQueryFilters field.
func (o *SLOWidgetDefinition) SetAdditionalQueryFilters(v string) {
	o.AdditionalQueryFilters = &v
}

// GetGlobalTimeTarget returns the GlobalTimeTarget field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetGlobalTimeTarget() string {
	if o == nil || o.GlobalTimeTarget == nil {
		var ret string
		return ret
	}
	return *o.GlobalTimeTarget
}

// GetGlobalTimeTargetOk returns a tuple with the GlobalTimeTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetGlobalTimeTargetOk() (*string, bool) {
	if o == nil || o.GlobalTimeTarget == nil {
		return nil, false
	}
	return o.GlobalTimeTarget, true
}

// HasGlobalTimeTarget returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasGlobalTimeTarget() bool {
	return o != nil && o.GlobalTimeTarget != nil
}

// SetGlobalTimeTarget gets a reference to the given string and assigns it to the GlobalTimeTarget field.
func (o *SLOWidgetDefinition) SetGlobalTimeTarget(v string) {
	o.GlobalTimeTarget = &v
}

// GetShowErrorBudget returns the ShowErrorBudget field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetShowErrorBudget() bool {
	if o == nil || o.ShowErrorBudget == nil {
		var ret bool
		return ret
	}
	return *o.ShowErrorBudget
}

// GetShowErrorBudgetOk returns a tuple with the ShowErrorBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetShowErrorBudgetOk() (*bool, bool) {
	if o == nil || o.ShowErrorBudget == nil {
		return nil, false
	}
	return o.ShowErrorBudget, true
}

// HasShowErrorBudget returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasShowErrorBudget() bool {
	return o != nil && o.ShowErrorBudget != nil
}

// SetShowErrorBudget gets a reference to the given bool and assigns it to the ShowErrorBudget field.
func (o *SLOWidgetDefinition) SetShowErrorBudget(v bool) {
	o.ShowErrorBudget = &v
}

// GetSloId returns the SloId field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetSloId() string {
	if o == nil || o.SloId == nil {
		var ret string
		return ret
	}
	return *o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetSloIdOk() (*string, bool) {
	if o == nil || o.SloId == nil {
		return nil, false
	}
	return o.SloId, true
}

// HasSloId returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasSloId() bool {
	return o != nil && o.SloId != nil
}

// SetSloId gets a reference to the given string and assigns it to the SloId field.
func (o *SLOWidgetDefinition) SetSloId(v string) {
	o.SloId = &v
}

// GetTimeWindows returns the TimeWindows field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTimeWindows() []WidgetTimeWindows {
	if o == nil || o.TimeWindows == nil {
		var ret []WidgetTimeWindows
		return ret
	}
	return o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTimeWindowsOk() (*[]WidgetTimeWindows, bool) {
	if o == nil || o.TimeWindows == nil {
		return nil, false
	}
	return &o.TimeWindows, true
}

// HasTimeWindows returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTimeWindows() bool {
	return o != nil && o.TimeWindows != nil
}

// SetTimeWindows gets a reference to the given []WidgetTimeWindows and assigns it to the TimeWindows field.
func (o *SLOWidgetDefinition) SetTimeWindows(v []WidgetTimeWindows) {
	o.TimeWindows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SLOWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *SLOWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *SLOWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *SLOWidgetDefinition) GetType() SLOWidgetDefinitionType {
	if o == nil {
		var ret SLOWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTypeOk() (*SLOWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SLOWidgetDefinition) SetType(v SLOWidgetDefinitionType) {
	o.Type = v
}

// GetViewMode returns the ViewMode field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetViewMode() WidgetViewMode {
	if o == nil || o.ViewMode == nil {
		var ret WidgetViewMode
		return ret
	}
	return *o.ViewMode
}

// GetViewModeOk returns a tuple with the ViewMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetViewModeOk() (*WidgetViewMode, bool) {
	if o == nil || o.ViewMode == nil {
		return nil, false
	}
	return o.ViewMode, true
}

// HasViewMode returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasViewMode() bool {
	return o != nil && o.ViewMode != nil
}

// SetViewMode gets a reference to the given WidgetViewMode and assigns it to the ViewMode field.
func (o *SLOWidgetDefinition) SetViewMode(v WidgetViewMode) {
	o.ViewMode = &v
}

// GetViewType returns the ViewType field value.
func (o *SLOWidgetDefinition) GetViewType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetViewTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewType, true
}

// SetViewType sets field value.
func (o *SLOWidgetDefinition) SetViewType(v string) {
	o.ViewType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AdditionalQueryFilters != nil {
		toSerialize["additional_query_filters"] = o.AdditionalQueryFilters
	}
	if o.GlobalTimeTarget != nil {
		toSerialize["global_time_target"] = o.GlobalTimeTarget
	}
	if o.ShowErrorBudget != nil {
		toSerialize["show_error_budget"] = o.ShowErrorBudget
	}
	if o.SloId != nil {
		toSerialize["slo_id"] = o.SloId
	}
	if o.TimeWindows != nil {
		toSerialize["time_windows"] = o.TimeWindows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type
	if o.ViewMode != nil {
		toSerialize["view_mode"] = o.ViewMode
	}
	toSerialize["view_type"] = o.ViewType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AdditionalQueryFilters *string                  `json:"additional_query_filters,omitempty"`
		GlobalTimeTarget       *string                  `json:"global_time_target,omitempty"`
		ShowErrorBudget        *bool                    `json:"show_error_budget,omitempty"`
		SloId                  *string                  `json:"slo_id,omitempty"`
		TimeWindows            []WidgetTimeWindows      `json:"time_windows,omitempty"`
		Title                  *string                  `json:"title,omitempty"`
		TitleAlign             *WidgetTextAlign         `json:"title_align,omitempty"`
		TitleSize              *string                  `json:"title_size,omitempty"`
		Type                   *SLOWidgetDefinitionType `json:"type"`
		ViewMode               *WidgetViewMode          `json:"view_mode,omitempty"`
		ViewType               *string                  `json:"view_type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ViewType == nil {
		return fmt.Errorf("required field view_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additional_query_filters", "global_time_target", "show_error_budget", "slo_id", "time_windows", "title", "title_align", "title_size", "type", "view_mode", "view_type"})
	} else {
		return err
	}
	if v := all.TitleAlign; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.ViewMode; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.AdditionalQueryFilters = all.AdditionalQueryFilters
	o.GlobalTimeTarget = all.GlobalTimeTarget
	o.ShowErrorBudget = all.ShowErrorBudget
	o.SloId = all.SloId
	o.TimeWindows = all.TimeWindows
	o.Title = all.Title
	o.TitleAlign = all.TitleAlign
	o.TitleSize = all.TitleSize
	o.Type = *all.Type
	o.ViewMode = all.ViewMode
	o.ViewType = *all.ViewType
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
