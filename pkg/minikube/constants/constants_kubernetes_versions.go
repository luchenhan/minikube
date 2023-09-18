/*
Copyright 2023 The Kubernetes Authors All rights reserved.

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

package constants

// This file is updated automatically, do not manually modify

// ValidKubernetesVersions is a list of Kubernetes versions in order from newest to oldest
// This is used when outputting Kubernetes versions and to select the latest patch version when unspecified
var ValidKubernetesVersions = []string{
	"v1.28.2",
	"v1.28.1",
	"v1.28.0",
	"v1.28.0-rc.1",
	"v1.28.0-rc.0",
	"v1.28.0-beta.0",
	"v1.28.0-alpha.4",
	"v1.28.0-alpha.3",
	"v1.28.0-alpha.2",
	"v1.28.0-alpha.1",
	"v1.27.6",
	"v1.27.5",
	"v1.27.4",
	"v1.27.3",
	"v1.27.2",
	"v1.27.1",
	"v1.27.0",
	"v1.27.0-rc.1",
	"v1.27.0-rc.0",
	"v1.27.0-beta.0",
	"v1.27.0-alpha.3",
	"v1.27.0-alpha.2",
	"v1.27.0-alpha.1",
	"v1.26.9",
	"v1.26.8",
	"v1.26.7",
	"v1.26.6",
	"v1.26.5",
	"v1.26.4",
	"v1.26.3",
	"v1.26.2",
	"v1.26.1",
	"v1.26.0",
	"v1.26.0-rc.1",
	"v1.26.0-rc.0",
	"v1.26.0-beta.0",
	"v1.26.0-alpha.3",
	"v1.26.0-alpha.2",
	"v1.26.0-alpha.1",
	"v1.25.14",
	"v1.25.13",
	"v1.25.12",
	"v1.25.11",
	"v1.25.10",
	"v1.25.9",
	"v1.25.8",
	"v1.25.7",
	"v1.25.6",
	"v1.25.5",
	"v1.25.4",
	"v1.25.3",
	"v1.25.2",
	"v1.25.1",
	"v1.25.0",
	"v1.25.0-rc.1",
	"v1.25.0-rc.0",
	"v1.25.0-beta.0",
	"v1.25.0-alpha.3",
	"v1.25.0-alpha.2",
	"v1.25.0-alpha.1",
	"v1.24.17",
	"v1.24.16",
	"v1.24.15",
	"v1.24.14",
	"v1.24.13",
	"v1.24.12",
	"v1.24.11",
	"v1.24.10",
	"v1.24.9",
	"v1.24.8",
	"v1.24.7",
	"v1.24.6",
	"v1.24.5",
	"v1.24.4",
	"v1.24.3",
	"v1.24.2",
	"v1.24.1",
	"v1.24.0",
	"v1.24.0-rc.1",
	"v1.24.0-rc.0",
	"v1.24.0-beta.0",
	"v1.24.0-alpha.4",
	"v1.24.0-alpha.3",
	"v1.24.0-alpha.2",
	"v1.24.0-alpha.1",
	"v1.23.17",
	"v1.23.16",
	"v1.23.15",
	"v1.23.14",
	"v1.23.13",
	"v1.23.12",
	"v1.23.11",
	"v1.23.10",
	"v1.23.9",
	"v1.23.8",
	"v1.23.7",
	"v1.23.6",
	"v1.23.5",
	"v1.23.4",
	"v1.23.3",
	"v1.23.2",
	"v1.23.1",
	"v1.23.0",
	"v1.23.0-rc.1",
	"v1.23.0-rc.0",
	"v1.23.0-beta.0",
	"v1.23.0-alpha.4",
	"v1.23.0-alpha.3",
	"v1.23.0-alpha.2",
	"v1.23.0-alpha.1",
	"v1.22.17",
	"v1.22.16",
	"v1.22.15",
	"v1.22.14",
	"v1.22.13",
	"v1.22.12",
	"v1.22.11",
	"v1.22.10",
	"v1.22.9",
	"v1.22.8",
	"v1.22.7",
	"v1.22.6",
	"v1.22.5",
	"v1.22.4",
	"v1.22.3",
	"v1.22.2",
	"v1.22.1",
	"v1.22.0",
	"v1.22.0-rc.0",
	"v1.22.0-beta.2",
	"v1.22.0-beta.1",
	"v1.22.0-beta.0",
	"v1.22.0-alpha.3",
	"v1.22.0-alpha.2",
	"v1.22.0-alpha.1",
	"v1.21.14",
	"v1.21.13",
	"v1.21.12",
	"v1.21.11",
	"v1.21.10",
	"v1.21.9",
	"v1.21.8",
	"v1.21.7",
	"v1.21.6",
	"v1.21.5",
	"v1.21.4",
	"v1.21.3",
	"v1.21.2",
	"v1.21.1",
	"v1.21.0",
	"v1.21.0-rc.0",
	"v1.21.0-beta.1",
	"v1.21.0-beta.0",
	"v1.21.0-alpha.3",
	"v1.21.0-alpha.2",
	"v1.21.0-alpha.1",
	"v1.20.15",
	"v1.20.14",
	"v1.20.13",
	"v1.20.12",
	"v1.20.11",
	"v1.20.10",
	"v1.20.9",
	"v1.20.8",
	"v1.20.7",
	"v1.20.6",
	"v1.20.5",
	"v1.20.4",
	"v1.20.3",
	"v1.20.2",
	"v1.20.1",
	"v1.20.0",
	"v1.20.0-rc.0",
	"v1.20.0-beta.2",
	"v1.20.0-beta.1",
	"v1.20.0-beta.0",
	"v1.20.0-alpha.3",
	"v1.20.0-alpha.2",
	"v1.20.0-alpha.1",
	"v1.19.16",
	"v1.19.15",
	"v1.19.14",
	"v1.19.13",
	"v1.19.12",
	"v1.19.11",
	"v1.19.10",
	"v1.19.9",
	"v1.19.8",
	"v1.19.7",
	"v1.19.6",
	"v1.19.5",
	"v1.19.4",
	"v1.19.3",
	"v1.19.2",
	"v1.19.1",
	"v1.19.0",
	"v1.19.0-rc.4",
	"v1.19.0-rc.3",
	"v1.19.0-rc.2",
	"v1.19.0-rc.1",
	"v1.19.0-beta.2",
	"v1.19.0-beta.1",
	"v1.19.0-beta.0",
	"v1.19.0-alpha.3",
	"v1.19.0-alpha.2",
	"v1.19.0-alpha.1",
	"v1.18.20",
	"v1.18.19",
	"v1.18.18",
	"v1.18.17",
	"v1.18.16",
	"v1.18.15",
	"v1.18.14",
	"v1.18.13",
	"v1.18.12",
	"v1.18.10",
	"v1.18.9",
	"v1.18.8",
	"v1.18.6",
	"v1.18.5",
	"v1.18.5-rc.1",
	"v1.18.4",
	"v1.18.3",
	"v1.18.2",
	"v1.18.1",
	"v1.18.0",
	"v1.18.0-rc.1",
	"v1.18.0-beta.2",
	"v1.18.0-beta.1",
	"v1.18.0-alpha.5",
	"v1.18.0-alpha.3",
	"v1.18.0-alpha.2",
	"v1.18.0-alpha.1",
	"v1.17.17",
	"v1.17.16",
	"v1.17.15",
	"v1.17.14",
	"v1.17.13",
	"v1.17.12",
	"v1.17.11",
	"v1.17.9",
	"v1.17.8",
	"v1.17.8-rc.1",
	"v1.17.7",
	"v1.17.6",
	"v1.17.5",
	"v1.17.4",
	"v1.17.3",
	"v1.17.2",
	"v1.17.1",
	"v1.17.0",
	"v1.17.0-rc.2",
	"v1.17.0-rc.1",
	"v1.17.0-beta.2",
	"v1.17.0-beta.1",
	"v1.17.0-alpha.3",
	"v1.17.0-alpha.2",
	"v1.17.0-alpha.1",
	"v1.16.15",
	"v1.16.14",
	"v1.16.13",
	"v1.16.12",
	"v1.16.12-rc.1",
	"v1.16.11",
	"v1.16.10",
	"v1.16.9",
	"v1.16.8",
	"v1.16.7",
	"v1.16.6",
	"v1.16.5",
	"v1.16.4",
	"v1.16.3",
	"v1.16.2",
	"v1.16.1",
	"v1.16.0",
	"v1.16.0-rc.2",
	"v1.16.0-rc.1",
	"v1.16.0-beta.2",
	"v1.16.0-beta.1",
	"v1.16.0-alpha.3",
	"v1.16.0-alpha.2",
	"v1.16.0-alpha.1",
	"v1.15.12",
	"v1.15.11",
	"v1.15.10",
	"v1.15.9",
	"v1.15.8",
	"v1.15.7",
	"v1.15.6",
	"v1.15.5",
	"v1.15.4",
	"v1.15.3",
	"v1.15.2",
	"v1.15.1",
	"v1.15.0",
	"v1.15.0-rc.1",
	"v1.15.0-beta.2",
	"v1.15.0-beta.1",
	"v1.15.0-alpha.3",
	"v1.15.0-alpha.2",
	"v1.15.0-alpha.1",
	"v1.14.10",
	"v1.14.9",
	"v1.14.8",
	"v1.14.7",
	"v1.14.6",
	"v1.14.5",
	"v1.14.4",
	"v1.14.3",
	"v1.14.2",
	"v1.14.1",
	"v1.14.0",
	"v1.14.0-rc.1",
	"v1.14.0-beta.2",
	"v1.14.0-beta.1",
	"v1.14.0-alpha.3",
	"v1.14.0-alpha.2",
	"v1.14.0-alpha.1",
	"v1.13.12",
	"v1.13.11",
	"v1.13.10",
	"v1.13.9",
	"v1.13.8",
	"v1.13.7",
	"v1.13.6",
	"v1.13.5",
	"v1.13.4",
	"v1.13.3",
	"v1.13.2",
	"v1.13.1",
	"v1.13.0",
	"v1.13.0-rc.2",
	"v1.13.0-rc.1",
	"v1.13.0-beta.2",
	"v1.13.0-beta.1",
	"v1.13.0-alpha.3",
	"v1.13.0-alpha.2",
	"v1.13.0-alpha.1",
	"v1.12.10",
	"v1.12.9",
	"v1.12.8",
	"v1.12.7",
	"v1.12.6",
	"v1.12.5",
	"v1.12.4",
	"v1.12.3",
	"v1.12.2",
	"v1.12.1",
	"v1.12.0",
	"v1.12.0-rc.2",
	"v1.12.0-rc.1",
	"v1.12.0-beta.2",
	"v1.12.0-beta.1",
	"v1.12.0-alpha.1",
	"v1.11.10",
	"v1.11.9",
	"v1.11.8",
	"v1.11.7",
	"v1.11.6",
	"v1.11.5",
	"v1.11.4",
	"v1.11.3",
	"v1.11.2",
	"v1.11.1",
	"v1.11.0",
	"v1.11.0-rc.3",
	"v1.11.0-rc.2",
	"v1.11.0-rc.1",
	"v1.11.0-beta.2",
	"v1.11.0-beta.1",
	"v1.11.0-alpha.2",
	"v1.11.0-alpha.1",
	"v1.10.13",
	"v1.10.12",
	"v1.10.11",
	"v1.10.10",
	"v1.10.9",
	"v1.10.8",
	"v1.10.7",
	"v1.10.6",
	"v1.10.5",
	"v1.10.4",
	"v1.10.3",
	"v1.10.2",
	"v1.10.1",
	"v1.10.0",
	"v1.10.0-rc.1",
	"v1.10.0-beta.4",
	"v1.10.0-beta.3",
	"v1.10.0-beta.2",
	"v1.10.0-beta.1",
	"v1.10.0-alpha.3",
	"v1.10.0-alpha.2",
	"v1.10.0-alpha.1",
	"v1.9.11",
	"v1.9.10",
	"v1.9.9",
	"v1.9.8",
	"v1.9.7",
	"v1.9.6",
	"v1.9.5",
	"v1.9.4",
	"v1.9.3",
	"v1.9.2",
	"v1.9.1",
	"v1.9.0",
	"v1.9.0-beta.2",
	"v1.9.0-beta.1",
	"v1.9.0-alpha.3",
	"v1.9.0-alpha.2",
	"v1.9.0-alpha.1",
	"v1.8.15",
	"v1.8.14",
	"v1.8.13",
	"v1.8.12",
	"v1.8.11",
	"v1.8.10",
	"v1.8.9",
	"v1.8.8",
	"v1.8.7",
	"v1.8.6",
	"v1.8.5",
	"v1.8.4",
	"v1.8.3",
	"v1.8.2",
	"v1.8.1",
	"v1.8.0",
	"v1.8.0-rc.1",
	"v1.8.0-beta.1",
	"v1.8.0-alpha.3",
	"v1.8.0-alpha.2",
	"v1.8.0-alpha.1",
	"v1.7.16",
	"v1.7.15",
	"v1.7.14",
	"v1.7.13",
	"v1.7.12",
	"v1.7.11",
	"v1.7.10",
	"v1.7.9",
	"v1.7.8",
	"v1.7.7",
	"v1.7.6",
	"v1.7.5",
	"v1.7.4",
	"v1.7.3",
	"v1.7.2",
	"v1.7.1",
	"v1.7.0",
	"v1.7.0-rc.1",
	"v1.7.0-beta.2",
	"v1.7.0-beta.1",
	"v1.7.0-alpha.4",
	"v1.7.0-alpha.3",
	"v1.7.0-alpha.2",
	"v1.7.0-alpha.1",
	"v1.6.13",
	"v1.6.12",
	"v1.6.11",
	"v1.6.10",
	"v1.6.9",
	"v1.6.8",
	"v1.6.7",
	"v1.6.6",
	"v1.6.5",
	"v1.6.4",
	"v1.6.3",
	"v1.6.2",
	"v1.6.1",
	"v1.6.0",
	"v1.6.0-rc.1",
	"v1.6.0-beta.4",
	"v1.6.0-beta.3",
	"v1.6.0-beta.2",
	"v1.6.0-beta.1",
	"v1.6.0-alpha.3",
	"v1.6.0-alpha.2",
	"v1.6.0-alpha.1",
	"v1.5.8",
	"v1.5.7",
	"v1.5.6",
	"v1.5.5",
	"v1.5.4",
	"v1.5.3",
	"v1.5.2",
	"v1.5.1",
	"v1.5.0",
	"v1.5.0-beta.3",
	"v1.5.0-beta.2",
	"v1.5.0-beta.1",
	"v1.5.0-alpha.2",
	"v1.5.0-alpha.1",
	"v1.4.12",
	"v1.4.9",
	"v1.4.8",
	"v1.4.7",
	"v1.4.6",
	"v1.4.5",
	"v1.4.4",
	"v1.4.3",
	"v1.4.2-beta.1",
	"v1.4.1",
	"v1.4.1-beta.2",
	"v1.4.0",
	"v1.4.0-beta.11",
	"v1.4.0-beta.10",
	"v1.4.0-beta.8",
	"v1.4.0-beta.7",
	"v1.4.0-beta.6",
	"v1.4.0-beta.5",
	"v1.4.0-beta.3",
	"v1.4.0-beta.2",
	"v1.4.0-beta.1",
	"v1.4.0-alpha.3",
	"v1.4.0-alpha.2",
	"v1.4.0-alpha.1",
	"v1.3.10",
	"v1.3.9",
	"v1.3.8",
	"v1.3.7",
	"v1.3.6",
	"v1.3.5",
	"v1.3.4",
	"v1.3.3",
	"v1.3.2",
	"v1.3.1",
	"v1.3.0",
	"v1.3.0-beta.3",
	"v1.3.0-beta.2",
	"v1.3.0-beta.1",
	"v1.3.0-alpha.5",
	"v1.3.0-alpha.4",
	"v1.3.0-alpha.3",
	"v1.3.0-alpha.2",
	"v1.3.0-alpha.1",
	"v1.2.7",
	"v1.2.6",
	"v1.2.5",
	"v1.2.4",
	"v1.2.3",
	"v1.2.2",
	"v1.2.1",
	"v1.2.0",
	"v1.2.0-beta.1",
	"v1.2.0-beta.0",
	"v1.2.0-alpha.8",
	"v1.2.0-alpha.7",
	"v1.2.0-alpha.6",
	"v1.2.0-alpha.5",
	"v1.2.0-alpha.4",
	"v1.2.0-alpha.3",
	"v1.2.0-alpha.2",
	"v1.2.0-alpha.1",
	"v1.1.8",
	"v1.1.7",
	"v1.1.4",
	"v1.1.3",
	"v1.1.2",
	"v1.1.1",
	"v1.1.1-beta.1",
	"v1.1.0-alpha.1",
	"v1.0.7",
	"v1.0.6",
	"v1.0.5",
	"v1.0.4",
	"v1.0.3",
	"v1.0.2",
	"v1.0.1",
	"v1.0.0",
	"v0.21.4",
	"v0.21.3",
	"v0.21.2",
	"v0.21.1",
	"v0.21.0",
	"v0.20.2",
	"v0.20.1",
	"v0.20.0",
	"v0.19.3",
	"v0.19.1",
	"v0.19.0",
	"v0.18.2",
	"v0.18.1",
	"v0.18.0",
	"v0.17.1",
	"v0.17.0",
	"v0.16.2",
	"v0.16.1",
	"v0.16.0",
	"v0.15.0",
	"v0.14.2",
	"v0.14.1",
	"v0.13.2",
	"v0.13.1",
	"v0.13.0",
	"v0.12.2",
	"v0.12.1",
	"v0.12.0",
	"v0.11.0",
	"v0.10.1",
	"v0.10.0",
	"v0.9.3",
	"v0.9.2",
	"v0.9.1",
	"v0.9.0",
	"v0.8.4",
	"v0.8.2",
	"v0.8.1",
	"v0.8.0",
	"v0.7.4",
	"v0.7.3",
	"v0.7.2",
	"v0.7.1",
	"v0.7.0",
	"v0.6.2",
	"v0.6.1",
	"v0.6.0",
	"v0.5.6",
	"v0.5.4",
	"v0.5.3",
	"v0.5.2",
	"v0.5.1",
	"v0.5",
	"v0.4.4",
	"v0.4.3",
	"v0.4.2",
	"v0.4.1",
	"v0.4",
}
