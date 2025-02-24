package dbmodels

type MarketMover struct {
	StockCode         string  `bson:"stock_symbol"`
	Company           string  `bson:"company_name"`
	TransactionsCount int32   `bson:"no_of_transactions"`
	HighestPrice      float64 `bson:"max_price"`
	LowestPrice       float64 `bson:"min_price"`
	OpeningPrice      float64 `bson:"opening_price"`
	ClosingPrice      float64 `bson:"closing_price"`
	Turnover          float64 `bson:"amount"`
	PreviousClose     float64 `bson:"previous_closing"`
	PriceChange       float64 `bson:"difference_rs"`
	PercentageChange  float64 `bson:"percent_change"`
	TradedVolume      int32   `bson:"volume"`
	TradeDate         string  `bson:"trade_date"`
}

type IPOAndFpoAlert struct {
	Ipo []IPOAlert `bson:"ipo"`
	Fpo []IPOAlert `bson:"fpo"`
}

type MarketMovers struct {
	Gainers []MarketMover `bson:"gainers"`
	Loser   []MarketMover `bson:"losers"`
}

type NepseIndex struct {
	MarketIndex          string  `bson:"index_name"`
	CurrentValue         float64 `bson:"index_value"`
	PreviousClose        float64 `bson:"previous_value"`
	OpeningValue         float64 `bson:"opening_value"`
	PercentageChange     float64 `bson:"percent_change"`
	PointChange          float64 `bson:"difference"`
	TotalTurnover        float64 `bson:"turnover"`
	TradedVolume         int32   `bson:"volume"`
	MarketCapitalization float64 `bson:"market_cap"`
	DailyHigh            float64 `bson:"day_high"`
	DailyLow             float64 `bson:"day_low"`
	YearlyHigh           float64 `bson:"year_high"`
	YearlyLow            float64 `bson:"year_low"`
	Date                 string  `bson:"as_of_date"`
}

type IPOAlert struct {
	IssuerName             string `bson:"company_name"`
	StockSymbol            string `bson:"stock_symbol"`
	ShareRegistrar         string `bson:"share_registrar"`
	IndustrySector         string `bson:"sector_name"`
	ShareType              string `bson:"share_type"`
	UnitPrice              string `bson:"price_per_unit"`
	Rating                 string `bson:"rating"`
	NumberOfShares         string `bson:"units"`
	MinimumUnits           string `bson:"min_units"`
	MaximumUnits           string `bson:"max_units"`
	TotalApplicationAmount string `bson:"total_amount"`
	ApplicationStartDateAD string `bson:"opening_date_ad"`
	ApplicationStartDateBS string `bson:"opening_date_bs"`
	ApplicationEndDateAD   string `bson:"closing_date_ad"`
	ApplicationEndDateBS   string `bson:"closing_date_bs"`
	ApplicationClosingTime string `bson:"closing_date_closing_time"`
	Status                 string `bson:"status"`
}

type Market struct {
	Symbol           string  `bson:"symbol"`
	Company          string  `bson:"company"`
	TradeVolume      int     `bson:"tradeVolume"`
	High             float64 `bson:"high"`
	Low              float64 `bson:"low"`
	Open             float64 `bson:"open"`
	Close            float64 `bson:"close"`
	TotalTradedValue float64 `bson:"totalTradedValue"`
	PrevClose        float64 `bson:"prevClose"`
	PriceChange      float64 `bson:"priceChange"`
	PercentChange    float64 `bson:"percentChange"`
	ShareVolume      int     `bson:"shareVolume"`
	LastUpdated      string  `bson:"lastupdated"`
}

type Indices struct {
	IndexName        string  `bson:"index_name"`
	IndexValue       float64 `bson:"index_value"`
	PreviousValue    float64 `bson:"previous_value"`
	OpeningValue     float64 `bson:"opening_value"`
	PercentChange    float64 `bson:"percent_change"`
	Difference       float64 `bson:"difference"`
	Turnover         float64 `bson:"turnover"`
	Volume           int32   `bson:"volume"`
	TotalCompanies   int32   `bson:"no_of_listed_companies"`
	TradedCompanies  int32   `bson:"no_of_traded_companies"`
	Transactions     int32   `bson:"no_of_transactions"`
	ListedShares     int64   `bson:"no_of_listed_shares"`
	MarketCap        float64 `bson:"market_cap"`
	DailyHigh        float64 `bson:"day_high"`
	DailyLow         float64 `bson:"day_low"`
	YearlyHigh       float64 `bson:"year_high"`
	YearlyLow        float64 `bson:"year_low"`
	ReportDate       string  `bson:"as_of_date"`
	ReportDateString string  `bson:"as_of_date_string"`
	GainingCompanies int32   `bson:"no_of_gainers"`
	LosingCompanies  int32   `bson:"no_of_losers"`
	Unchanged        int32   `bson:"no_of_unchanged"`
}

type MarketStatus struct {
	IsOpen string `bson:"isOpen"`
}
