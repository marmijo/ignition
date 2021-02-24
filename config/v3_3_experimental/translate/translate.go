// Copyright 2020 Red Hat, Inc.
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

package translate

import (
	"github.com/coreos/ignition/v2/config/translate"
	old_types "github.com/coreos/ignition/v2/config/v3_2/types"
	"github.com/coreos/ignition/v2/config/v3_3_experimental/types"
)

func translateIgnition(old old_types.Ignition) (ret types.Ignition) {
	// use a new translator so we don't recurse infinitely
	translate.NewTranslator().Translate(&old, &ret)
	ret.Version = types.MaxVersion.String()
	return
}

func translateClevis(old old_types.Clevis) (ret types.Clevis) {
	tr := translate.NewTranslator()
	tr.AddCustomTranslator(translateClevisCustom)
	if old.Custom != nil {
		tr.Translate(old.Custom, &ret.Custom)
	}
	tr.Translate(&old.Tang, &ret.Tang)
	tr.Translate(&old.Threshold, &ret.Threshold)
	tr.Translate(&old.Tpm2, &ret.Tpm2)
	return
}

func translateClevisCustom(old old_types.Custom) (ret types.ClevisCustom) {
	tr := translate.NewTranslator()
	tr.Translate(&old.Config, &ret.Config)
	tr.Translate(&old.NeedsNetwork, &ret.NeedsNetwork)
	tr.Translate(&old.Pin, &ret.Pin)
	return
}

func Translate(old old_types.Config) (ret types.Config) {
	tr := translate.NewTranslator()
	tr.AddCustomTranslator(translateIgnition)
	tr.AddCustomTranslator(translateClevis)
	tr.Translate(&old, &ret)
	return
}
