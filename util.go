package bocBankSDK

import (
	"errors"
	"os"
	"strings"

	"github.com/jointownchain/boc_bank_sdk/Modules"
	ecode "github.com/jointownchain/mic_common/exception"
)

// 从环境变量获取 银行的 URL
func GetBankUrl() (string, error) {
	url := os.Getenv(Modules.BOCB2E_BANK_URL_ENV)
	if url == "" {
		return "", errors.New("BOCB2E_BANK_URL_ENV not set!")
	}
	return url, nil
}

func IsRspCodeOK(code string) bool {
	if code != Modules.BOCB2E_RESPONSE_CODE_OK && code != Modules.BOCB2E_REPONSE_CODE_TBC {
		return false
	}
	return true
}

func IsBankError(err error) bool {
	if err == ecode.RequestErr {
		return true
	}
	return strings.Contains(err.Error(), "BankSDKError")
}
