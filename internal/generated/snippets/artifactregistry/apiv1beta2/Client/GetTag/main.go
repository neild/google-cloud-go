// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START artifactregistry_generated_artifactregistry_apiv1beta2_Client_GetTag]

package main

import (
	"context"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1beta2"
	artifactregistrypb "google.golang.org/genproto/googleapis/devtools/artifactregistry/v1beta2"
)

func main() {
	// import artifactregistrypb "google.golang.org/genproto/googleapis/devtools/artifactregistry/v1beta2"

	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &artifactregistrypb.GetTagRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END artifactregistry_generated_artifactregistry_apiv1beta2_Client_GetTag]