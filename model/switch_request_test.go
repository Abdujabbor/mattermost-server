// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwitchRequestJson(t *testing.T) {
	o := SwitchRequest{Email: NewId(), Password: NewId()}
	json := o.ToJson()
	ro := SwitchRequestFromJson(strings.NewReader(json))

	require.Equal(t, o.Email, ro.Email, "Emails do not match")
}
