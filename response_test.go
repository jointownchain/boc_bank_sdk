package bocBankSDK

import (
	"encoding/json"
	"fmt"
	"gitlab.koblitzdigital.com/lib/bocBankSDK/Modules"
	"reflect"
	"testing"
)

func TestParseXML(t *testing.T) {
	rspXML := `<?xml version="1.0" encoding="UTF-8"?>
<bocb2e version="120" security="true" lang="zh_CN">
   <head>
      <termid>1</termid>
      <trnid>2</trnid>
      <custid>3</custid>
      <cusopr>4</cusopr>
      <trncod>5</trncod>
      <token>6</token>
   </head>
   <trans>
	   <trn-b2e0009-rs>
		  <status>
			 <rspcod>123</rspcod>
			 <rspmsg>123123</rspmsg>
		  </status>
		  <b2e0009-rs>
			 <status>
				<rspcod>123</rspcod>
				<rspmsg>123123</rspmsg>
			 </status>
			 <insid>123123</insid>
			 <obssid>123123</obssid>
		  </b2e0009-rs>
	   </trn-b2e0009-rs>
	</trans>
</bocb2e>
`

	in, err := parseResponseXML([]byte(rspXML), Modules.BOCB2E_OPTS_HOSPITAL_PAYMENT)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(in))
	rsp := in.(*Modules.HospitalPaymentRspTrans)
	rspBytes, err := json.Marshal(rsp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rspBytes))
}
