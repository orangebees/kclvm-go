package kpm

import (
	"github.com/orangebees/go-oneutils/GlobalStore"
)

type Require struct {
	RequireBase
	//别名
	Alias string `json:"alias,omitempty"`
}
type RequirePlus struct {
	//引用计数
	Count int `json:"count"`
	//校验和 sha512
	Integrity GlobalStore.Integrity `json:"integrity"`
}

type RequireBase struct {
	////包类型 git，registry
	//Type string `json:"type"`
	////包名，确定包的命名空间
	//Name string `json:"name"`
	////确定此包的版本
	//Version PkgVersion `json:"version"`

	RequirePkgString
	//校验和 sha512
	Integrity GlobalStore.Integrity `json:"integrity"`

	//git:github.com/a/b@v0.0.1
	//git:github.com/a/b@v0.0.0#asdfghjkl
	//reg:b@v0.0.1
}

type RequirePkgString struct {
	//包类型 git，registry
	Type string `json:"type"`
	//包名，确定包的命名空间
	Name string `json:"name"`
	//确定此包的版本
	Version PkgVersion `json:"version"`

	//git:github.com/a/b@v0.0.1
	//git:github.com/a/b@v0.0.0#asdfghjkl
	//reg:b@v0.0.1
}

func (rps *RequirePkgString) GetPkgString() PkgString {
	return PkgString(rps.Type + ":" + rps.Name + "@" + string(rps.Version))
}
