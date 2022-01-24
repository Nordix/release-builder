// Copyright Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

import (
	"fmt"
	"path"

	"istio.io/pkg/log"
	"istio.io/release-builder/pkg/model"
	"istio.io/release-builder/pkg/util"
)

// Sbom generates Software Bill Of Materials for istio repo in an SPDX readable format.
func GenerateBillOfMaterials(manifest model.Manifest) error {
	// Retrieve istio repository path to run the sbom generator
	istioRepoDir := manifest.RepoDir("istio")
	sourceSbomFile := path.Join(manifest.OutDir(), "istio-source.spdx")
	sourceSbomNamespace := fmt.Sprintf("https://storage.googleapis.com/istio-release/releases/%s/istio-source.spdx",
		manifest.Version)
	releaseSbomFile := path.Join(manifest.OutDir(), "istio-release.spdx")
	releaseSbomNamespace := fmt.Sprintf("https://storage.googleapis.com/istio-release/releases/%s/istio-release.spdx",
		manifest.Version)

	// Run bom generator to generate the software bill of materials(SBOM) for istio.
	log.Infof("Generating Software Bill of Materials for istio release artifacts")
	if err := util.VerboseCommand("bom", "generate", "--namespace", releaseSbomNamespace,
		"--dirs", manifest.OutDir(), "--ignore", "licenses,'*.sha256'",
		"--output", releaseSbomFile).Run(); err != nil {
		return fmt.Errorf("couldn't generate sbom for istio release artifacts: %v", err)
	}

	// Run bom generator to generate the software bill of materials(SBOM) for istio.
	log.Infof("Generating Software Bill of Materials for istio source code")
	if err := util.VerboseCommand("bom", "generate", "--namespace", sourceSbomNamespace, "-d", istioRepoDir, "--output", sourceSbomFile).Run(); err != nil {
		return fmt.Errorf("couldn't generate sbom for istio source: %v", err)
	}
	return nil
}
