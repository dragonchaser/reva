// Copyright 2018-2021 CERN
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
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package loader

import (
	// Load cbox specific drivers.
	_ "github.com/cs3org/reva/v2/pkg/cbox/favorite/sql"
	_ "github.com/cs3org/reva/v2/pkg/cbox/group/rest"
	_ "github.com/cs3org/reva/v2/pkg/cbox/publicshare/sql"
	_ "github.com/cs3org/reva/v2/pkg/cbox/share/sql"
	_ "github.com/cs3org/reva/v2/pkg/cbox/storage/eoshomewrapper"
	_ "github.com/cs3org/reva/v2/pkg/cbox/storage/eoswrapper"
	_ "github.com/cs3org/reva/v2/pkg/cbox/user/rest"
)
