package main

//go:generate go run ca/database/migration/generate/generate.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/roosr/gotest/ca/csr"
	"github.com/roosr/gotest/ca/database"
	caErr "github.com/roosr/gotest/ca/error"
	"github.com/roosr/gotest/ca/http"
	"github.com/roosr/gotest/ca/test"
	"github.com/roosr/gotest/util"
)

type Rick struct {
	name string
}

type Res struct {
	Name string
}

func (v *Rick) CtHandler(w http.ResponseWriter, r *http.Request) {

	res := Res{
		Name: dchttp.GetPathParam(r, "name2"),
	}

	dchttp.WriteJsonResponse(w, res)
}

func main() {

	e1, e2 := test.GetA()
	if e1 == e2 {
		return
	}

	cadb.RunSchemaMigration()

	dao := cadb.New()
	user, err1 := dao.GetName("id2")
	if err1 != nil {
		if err1.Error() == caErr.NotFound {
			log.Printf("%s", "bad id")
		}

		log.Printf("%s", err1.Error())
	}
	log.Printf("%s", user.Username)

	v := &Rick{
		name: "rick",
	}

	server := dchttp.New("0.0.0.0", 3333)
	server.AddGetRoute("/test/{name}", v.CtHandler)
	go server.Run()
	util.WaitForStopSignal()
	server.Stop()

	log.SetFlags(log.LUTC | log.LstdFlags | log.Llongfile)
	log.Printf("error name")
	glog.Info("event=glog")

	csr, err := csr.ParseCsr("MIICijCCAXICAQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJ6tZ0djuKkuNDfDxj+7aIrtOaxIuyuweVwj1XGgjygY+mMVoanr1BsQfNf5rvqaifiUGNZo6UABx/GdoTPBPacU9lflKwcPNLuEO1nDz7Pwy7HmIsyAtqB24Yp62eyTmIEH+CdKXSSDVU36i9SS46CoHEpIeZeQ+8KfLTFYWVCCHOwL6NTemq/FrSynb3ygrG2HYFDMtSoMtOnhjYxdNxiu+JgoObu1diKuzsclHFwdb5SjQ+NA/UnVBs4FexnYohm8QnmLi0JdlgWaaXZbpNDe1B+ldjrEPBXq2GRbkSv6kdCnYPXNU9DWtWkJi+Oyi6IhCYngJCQC4Avgmy9s7Q8CAwEAAaAAMA0GCSqGSIb3DQEBCwUAA4IBAQAogbMdd7eoXJ8PfIdIjeJymIPCPKPVTwVCUkOa81SHKSkv/sBj059IeaNbBMHNjklWyiQd82Yto5VuyvQaw15ev44GGsDhdmmZJgT5rGb/F6pj5WNlN/WlLJMIf7OBim595I5KW93kKiYLoKPfCZN5Uxg+mZGcMmI3kyuSXlAf452ad3XMzgcll/S5uEXwDoshPYWa23Q0JZX7rTXb4oo0iG7+vxhrHy94HLAUcrUvM5OS83HZqbv34nenG+XZEU6HelJ/SS3z5yQBFbO7cennrp2IoquYANs83CwI4cm3OiXDQ4gfApf3htiztM44S7SXmwalLXVJmvSiwa7oV79j")
	//csr, err := ca.ParseCsr("MIICizCCAXMCAQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDCCASMwDQYJKoZIhvcNAQEBBQADggEQADCCAQsCggEBAJ6tZ0djuKkuNDfDxj+7aIrtOaxIuyuweVwj1XGgjygY+mMVoanr1BsQfNf5rvqaifiUGNZo6UABx/GdoTPBPacU9lflKwcPNLuEO1nDz7Pwy7HmIsyAtqB24Yp62eyTmIEH+CdKXSSDVU36i9SS46CoHEpIeZeQ+8KfLTFYWVCCHOwL6NTemq/FrSynb3ygrG2HYFDMtSoMtOnhjYxdNxiu+JgoObu1diKuzsclHFwdb5SjQ+NA/UnVBs4FexnYohm8QnmLi0JdlgWaaXZbpNDe1B+ldjrEPBXq2GRbkSv6kdCnYPXNU9DWtWkJi+Oyi6IhCYngJCQC4Avgmy9s7Q8CBAABAAGgADANBgkqhkiG9w0BAQsFAAOCAQEAKIGzHXe3qFyfD3yHSI3icpiDwjyj1U8FQlJDmvNUhykpL/7AY9OfSHmjWwTBzY5JVsokHfNmLaOVbsr0GsNeXr+OBhrA4XZpmSYE+axm/xeqY+VjZTf1pSyTCH+zgYpufeSOSlvd5ComC6Cj3wmTeVMYPpmRnDJiN5Mrkl5QH+Odmnd1zM4HJZf0ubhF8A6LIT2Fmtt0NCWV+6012+KKNIhu/r8Yax8veBywFHK1LzOTkvNx2am79+J3pxvl2RFOh3pSf0kt8+ckARWzu3Hp566diKKrmADbPNwsCOHJtzolw0OIHwKX94bYs7TOOEu0l5sGpS11SZr0osGu6Fe/Yw==")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%s", csr.CommonName)

}
