package mock

import (
	"encoding/xml"
	"os"

	"github.com/jointownchain/boc_bank_sdk/Modules"
)

// only for tests

func MockSetBankURL() {
	// _ = os.Setenv(Modules.BOCB2E_BANK_URL_ENV, "http://192.168.14.1:8080/B2EC/E2BServlet")
	_ = os.Setenv(Modules.BOCB2E_BANK_URL_ENV, "http://58.215.120.156:38080/B2EC/E2BServlet")
}

func MockXMLHead(trncod string) *Modules.Head {
	MockSetBankURL()
	head := &Modules.Head{xml.Name{}, "", "E192168202063",
		"04051003128",
		"133832424",
		"136170292",
		trncod,
		"0D6886D4498DB958C471131A189DF9CF"}
	return head
}

//  <termid>E192168202063</termid>
//  <trnid>04051003128</trnid>
//  <custid>133832424</custid>
//  <cusopr>136170292</cusopr>
//  <trncod>b2e0005</trncod>
//  <token>0D6886D4498DB958C471131A189DF9CF</token>

// 集团号 133736208
// 操作员号 136009152 (52， 53， 56， 57)
// 账号

func MockCompanyXMLHead(trncod string) *Modules.Head {
	MockSetBankURL()
	head := &Modules.Head{xml.Name{}, "", "E192168202063",
		"04051003128",
		"133832427",
		"136170292",
		trncod,
		"0D6886D4498DB958C471131A189DF9CF"}
	return head
}

func MockFactoringXMLHead(trncod string) *Modules.Head {
	MockSetBankURL()
	head := &Modules.Head{xml.Name{}, "", "E192168202063",
		"04051003128",
		"133832427",
		"136170403",
		trncod,
		"0D6886D4498DB958C471131A189DF9CF"}
	return head
}

func MockFactoring0312XMLHead(trncod string) *Modules.Head {
	MockSetBankURL()
	head := &Modules.Head{xml.Name{}, "", "E192168202063",
		"04051003128",
		"133832423",
		"136170292",
		trncod,
		"0D6886D4498DB958C471131A189DF9CF"}
	return head
}
