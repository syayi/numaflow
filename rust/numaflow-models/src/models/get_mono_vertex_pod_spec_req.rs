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

#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GetMonoVertexPodSpecReq {
    #[serde(rename = "DefaultResources")]
    pub default_resources: k8s_openapi::api::core::v1::ResourceRequirements,
    #[serde(rename = "Env")]
    pub env: Vec<k8s_openapi::api::core::v1::EnvVar>,
    #[serde(rename = "Image")]
    pub image: String,
    #[serde(rename = "PullPolicy")]
    pub pull_policy: String,
}

impl GetMonoVertexPodSpecReq {
    pub fn new(
        default_resources: k8s_openapi::api::core::v1::ResourceRequirements,
        env: Vec<k8s_openapi::api::core::v1::EnvVar>,
        image: String,
        pull_policy: String,
    ) -> GetMonoVertexPodSpecReq {
        GetMonoVertexPodSpecReq {
            default_resources,
            env,
            image,
            pull_policy,
        }
    }
}
