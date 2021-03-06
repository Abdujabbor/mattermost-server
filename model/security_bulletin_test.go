// Copyright (c) 2017 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSecurityBulletinToFromJson(t *testing.T) {
	b := SecurityBulletin{
		Id:               NewId(),
		AppliesToVersion: NewId(),
	}

	j := b.ToJson()
	b1 := SecurityBulletinFromJson(strings.NewReader(j))

	require.Equal(t, b, *b1)

	// Malformed JSON
	s2 := `{"wat"`
	b2 := SecurityBulletinFromJson(strings.NewReader(s2))
	require.Nil(t, b2)
}

func TestSecurityBulletinsToFromJson(t *testing.T) {
	b := SecurityBulletins{
		{
			Id:               NewId(),
			AppliesToVersion: NewId(),
		},
		{
			Id:               NewId(),
			AppliesToVersion: NewId(),
		},
	}

	j := b.ToJson()

	b1 := SecurityBulletinsFromJson(strings.NewReader(j))

	require.Len(t, b1, 2)

	// Malformed JSON
	s2 := `{"wat"`
	b2 := SecurityBulletinsFromJson(strings.NewReader(s2))

	require.Len(t, b2, 0)
}
