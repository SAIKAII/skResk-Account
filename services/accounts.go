package services

import (
	"time"

	"github.com/SAIKAII/skResk-Infra/base"

	"github.com/shopspring/decimal"
)

var IAccountService AccountService

func GetAccountService() AccountService {
	base.Check(IAccountService)
	return IAccountService
}

type AccountService interface {
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)
	Transfer(dto AccountTransferDTO) (TransferedStatus, error)
	StoreValue(dto AccountTransferDTO) (TransferedStatus, error)
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
	GetAccount(accountNo string) *AccountDTO
}

// 账户交易的参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

// 账户转账
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	Amount      decimal.Decimal
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}

// 账户创建
type AccountCreatedDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  int
	CurrencyCode string
	Amount       string
}

// 账户信息
type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
	Balance   decimal.Decimal
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 账户流水
type AccountLogDTO struct {
	LogNo           string          // 流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string          // 交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string          // 账户编号 账户ID
	TargetAccountNo string          // 账户编号 目标账户ID
	UserId          string          // 用户编号
	Username        string          // 用户名称
	TargetUserId    string          // 目标用户编号
	TargetUsername  string          // 目标用户名称
	Amount          decimal.Decimal // 交易金额，该交易涉及的金额
	Balance         decimal.Decimal // 交易后余额，该交易后的金额
	ChangeType      ChangeType      // 流水交易类型：0创建账户 >0收入类型 <0支出类型
	ChangeFlag      ChangeFlag      // 交易变化标识：-1出帐 1进账
	Status          int             // 交易状态
	Desc            string          // 交易描述
	CreatedAt       time.Time       // 创建时间
}
