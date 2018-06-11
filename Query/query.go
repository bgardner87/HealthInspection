package query

type Results []QueryResult

type Query interface {
	/*
	 * TODO
	 * Add additional fields to the query, beyond these basics.
	 */
	QueryWithBusinessName(name string) Query
	GetBusinessName() string
	QueryWithAddress(address string) Query
	QueryWithCity(city string) Query
	QueryWithState(state string) Query
	Print()
	Execute() (*Results, error)
}

type QueryResult interface {
	Print()
}
