/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Openapi Generator. DO NOT EDIT.

/// FixedWindow : FixedWindow describes a fixed window

#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct FixedWindow {
    #[serde(rename = "length", skip_serializing_if = "Option::is_none")]
    pub length: Option<kube::core::Duration>,
    /// Streaming should be set to true if the reduce udf is streaming.
    #[serde(rename = "streaming", skip_serializing_if = "Option::is_none")]
    pub streaming: Option<bool>,
}

impl FixedWindow {
    /// FixedWindow describes a fixed window
    pub fn new() -> FixedWindow {
        FixedWindow {
            length: None,
            streaming: None,
        }
    }
}
