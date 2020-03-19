// +build apitests

/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package tests

import "testing"

func Test_Terminators(t *testing.T) {
	ctx := NewTestContext(t)
	defer ctx.teardown()
	ctx.startServer()
	ctx.requireAdminLogin()

	service := ctx.AdminSession.requireNewService(nil, nil)
	edgeRouter := ctx.createAndEnroll()
	terminator := ctx.AdminSession.requireNewTerminator(service.id, edgeRouter.id, "transport", "tcp:localhost:2020")

	ctx.AdminSession.validateEntityWithLookup(terminator)
	ctx.AdminSession.validateEntityWithQuery(terminator)
	ctx.AdminSession.validateAssociations(service, terminator.getEntityType(), terminator)
}
