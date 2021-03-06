// automatically generated by the FlatBuffers compiler, do not modify

package fb

const (
	///< No Halt Type
	HaltTypenone = 0
	///< Unspecified news-related halt
	HaltTypenews = 1
	///< Denotes a regulatory trading halt when relevant news influencing the security is being disseminated. Trading is suspended until the primary market determines that an adequate publication or disclosure of information has occurred.
	HaltTypenews_disseminated = 2
	///< Denotes a non-regulatory halt condition where there is a significant imbalance of buy or sell orders.
	HaltTypeorder_imbalance = 3
	///< Denotes a non-regulatory Trading Halt. The ability to trade a security by a Participant is temporarily inhibited due to a systems, equipment or communications facility problem or for other technical reasons.
	HaltTypeequipment_change = 4
	///< Unspecified halt requiring additional information before resuming trading.
	HaltTypepending_additional_info = 5
	///< A financial status designation used to denote the ability to create new shares of this Exchange Traded Product (ETP) has been temporarily suspended by the ETP Issuer. ETPs that are closed for Creations typically are allowed to continue trading on the listing market once the ETP Issuer publishes the press release.
	HaltTypesuspended = 6
	///< Denotes a regulatory Trading halt mandated by the SEC for this security.
	HaltTypesec = 7
	///< Unspecified halt.
	HaltTypenot_specified = 8
	///< Denotes a five-minute regulatory trading halt (pause) for an individual security that does not exit a Limit State within 15 seconds
	HaltTypeluld_pause = 9
	///< The level 1 market-wide circuit breaker (MWCB) has been triggered due to a 7%+ decline in S&P500 from last-session close. If before 3:25PM, all securities are halted for 15 min. If at or after 3:25PM trading continues unless there is a Level 3 MWCB.
	HaltTypemarketwide_halt_level1 = 10
	///< the level 2 market-wide circuit breaker (MWCB) has been triggered due to a 13% decline in S&P500 from last-session close. If before 3:25PM all securities are halted for 15 min. If after 3:25PM, trading continues unless there is a Level 3 MWCB.
	HaltTypemarketwide_halt_level2 = 11
	///< The level 3 market-wide circuit breaker (MWCB) has been triggered due to a 20% decline in S&P500 from last-session close. All equities are halted for the remainder of the day.
	HaltTypemarketwide_halt_level3 = 12
	///< Indicates the deactivation of a market-wide circuit breaker. This should only occur for level 1 and level 2 MWCB events.
	HaltTypemarketwide_halt_resume = 13
	///< Denotes a five-minute regulatory trading halt (pause) for an individual security that does not exit a Limit State within 15 seconds. The limit-state is calculated depending on the exchange. This is 5% for >$3.00 S&P 500, Russel 1000 securities, and certain ETPs, 10% for all other securities > $3.00. See: http://cdn.batstrading.com/resources/membership/BATS_US_Equities_Limit_Up_Limit_Down_FAQ.pdf
	HaltTypeluld_straddle = 14
	///< Halt due to unusual market activity. (Note: Find CTA Multicast equivalent)
	HaltTypeextraordinary_market_activity = 15
	///< Indicates an unspecified halt for an exchange traded product
	HaltTypeetf = 16
	///< Indicates a halt issued by an exchange for failure to meet listing or other unspecified regulatory requirements
	HaltTypenon_compliance = 17
	///< A regulatory halt issued for equities not meeting reporting requirements
	HaltTypefilings_not_current = 18
	///< Halt reason issued for exchange operations being impacted. For instance, a storm.
	HaltTypeoperations = 19
	///< Pseudo-halt generated for IPO not occurring at market open
	HaltTypeipo_pending = 20
	///< Halted due to an intra-day event like a split. Rare.
	HaltTypecorporate_action = 21
	///< Quotations have temporarily become unavailable for an unspecified reason.
	HaltTypequote_unavailable = 22
	///< Generic halt condition for a single stock
	HaltTypesingle_stock_pause = 23
	///< Generic resume condition for a single stock
	HaltTypesingle_stock_pause_resume = 24
)

var EnumNamesHaltType = map[int]string{
	HaltTypenone:                          "none",
	HaltTypenews:                          "news",
	HaltTypenews_disseminated:             "news_disseminated",
	HaltTypeorder_imbalance:               "order_imbalance",
	HaltTypeequipment_change:              "equipment_change",
	HaltTypepending_additional_info:       "pending_additional_info",
	HaltTypesuspended:                     "suspended",
	HaltTypesec:                           "sec",
	HaltTypenot_specified:                 "not_specified",
	HaltTypeluld_pause:                    "luld_pause",
	HaltTypemarketwide_halt_level1:        "marketwide_halt_level1",
	HaltTypemarketwide_halt_level2:        "marketwide_halt_level2",
	HaltTypemarketwide_halt_level3:        "marketwide_halt_level3",
	HaltTypemarketwide_halt_resume:        "marketwide_halt_resume",
	HaltTypeluld_straddle:                 "luld_straddle",
	HaltTypeextraordinary_market_activity: "extraordinary_market_activity",
	HaltTypeetf:                           "etf",
	HaltTypenon_compliance:                "non_compliance",
	HaltTypefilings_not_current:           "filings_not_current",
	HaltTypeoperations:                    "operations",
	HaltTypeipo_pending:                   "ipo_pending",
	HaltTypecorporate_action:              "corporate_action",
	HaltTypequote_unavailable:             "quote_unavailable",
	HaltTypesingle_stock_pause:            "single_stock_pause",
	HaltTypesingle_stock_pause_resume:     "single_stock_pause_resume",
}
