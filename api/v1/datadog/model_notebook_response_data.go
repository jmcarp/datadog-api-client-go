/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// NotebookResponseData The data for a notebook.
type NotebookResponseData struct {
	Attributes NotebookResponseDataAttributes `json:"attributes"`
	// Unique notebook ID, assigned when you create the notebook.
	Id   int64                `json:"id"`
	Type NotebookResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewNotebookResponseData instantiates a new NotebookResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookResponseData(attributes NotebookResponseDataAttributes, id int64, type_ NotebookResourceType) *NotebookResponseData {
	this := NotebookResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = type_
	return &this
}

// NewNotebookResponseDataWithDefaults instantiates a new NotebookResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookResponseDataWithDefaults() *NotebookResponseData {
	this := NotebookResponseData{}
	var type_ NotebookResourceType = NOTEBOOKRESOURCETYPE_NOTEBOOKS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *NotebookResponseData) GetAttributes() NotebookResponseDataAttributes {
	if o == nil {
		var ret NotebookResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *NotebookResponseData) GetAttributesOk() (*NotebookResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *NotebookResponseData) SetAttributes(v NotebookResponseDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value
func (o *NotebookResponseData) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotebookResponseData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotebookResponseData) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value
func (o *NotebookResponseData) GetType() NotebookResourceType {
	if o == nil {
		var ret NotebookResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookResponseData) GetTypeOk() (*NotebookResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotebookResponseData) SetType(v NotebookResourceType) {
	o.Type = v
}

func (o NotebookResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *NotebookResponseData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *NotebookResponseDataAttributes `json:"attributes"`
		Id         *int64                          `json:"id"`
		Type       *NotebookResourceType           `json:"type"`
	}{}
	all := struct {
		Attributes NotebookResponseDataAttributes `json:"attributes"`
		Id         int64                          `json:"id"`
		Type       NotebookResourceType           `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
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
	o.Attributes = all.Attributes
	o.Id = all.Id
	o.Type = all.Type
	return nil
}
