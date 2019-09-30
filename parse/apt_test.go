// Copyright 2019 Sonatype Inc.
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
package parse

import (
	"testing"
)

func TestParseAptList(t *testing.T) {
	var list []string
	list = append(list, "libedit2 3.1-20170329-1")
	list = append(list, "libmount1 2.31.1-0.4ubuntu3.3")
	list = append(list, "zlib1g 1:1.2.11.dfsg-0ubuntu2")
	result := ParseAptList(list)

	if len(result.Projects) != 3 {
		t.Errorf("Didn't work")
	}

	if result.Projects[0].Name != "libedit2" || result.Projects[0].Version != "3.1" {
		t.Errorf("Libedit dep did not match result. Actual %s", result.Projects[0])
	}
	if result.Projects[1].Name != "libmount1" || result.Projects[1].Version != "2.31.1" {
		t.Errorf("Libmount dep did not match result. Actual %s", result.Projects[1])
	}
	if result.Projects[2].Name != "zlib1g" || result.Projects[2].Version != "1.2.11" {
		t.Errorf("Zlib1g dep did not match result. Actual %s", result.Projects[2])
	}
}

func TestParseAptListFromStdIn(t *testing.T) {
	var list []string
	list = append(list, "Listing...")
	list = append(list, "findutils/bionic,now 4.6.0+git+20170828-2 amd64 [installed]")
	list = append(list, "debconf/bionic-updates,now 1.5.66ubuntu1 all [installed]")
	list = append(list, "base-files/bionic-updates,now 10.1ubuntu2.6 amd64 [installed]")
	list = append(list, "Done")
	result := ParseAptListFromStdIn(list)

	if len(result.Projects) != 3 {
		t.Errorf("Didn't work")
	}

	if result.Projects[0].Name != "findutils" || result.Projects[0].Version != "4.6.0" {
		t.Errorf("findutils dep did not match result. Actual %s", result.Projects[0])
	}
	if result.Projects[1].Name != "debconf" || result.Projects[1].Version != "1.5.66" {
		t.Errorf("debconf dep did not match result. Actual %s", result.Projects[1])
	}
	if result.Projects[2].Name != "base-files" || result.Projects[2].Version != "10.1" {
		t.Errorf("base-files dep did not match result. Actual %s", result.Projects[2])
	}
}

func TestParseAptCanDealWithNonSemVerAndWillJustSkipIt(t *testing.T) {
	var list []string
	list = append(list, "Listing...")
	list = append(list, "ca-certificates/bionic,bionic-updates,now 20180409 all [installed,automatic]")
	list = append(list, "Done")
	result := ParseAptListFromStdIn(list)

	if len(result.Projects) != 0 {
		t.Errorf("Didn't work")
	}
}
