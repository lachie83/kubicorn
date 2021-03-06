// Copyright 2017 Google Inc. All Rights Reserved.
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

package profiler_test

import (
	"cloud.google.com/go/profiler"
)

func ExampleStart() {
	// The caller should provide the target string in the config so Cloud
	// Profiler knows how to group the profile data. Otherwise the target
	// string is set to "unknown".
	//
	// Optionally DebugLogging can be set in the config to enable detailed
	// logging from profiler.
	err := profiler.Start(&profiler.Config{Target: "my-target"})
	if err != nil {
		//TODO: Handle error.
	}
}
