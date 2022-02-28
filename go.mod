module app

go 1.17

replace github.com/digitalex70/celeritas => ../celeritas

require (
	github.com/CloudyKit/jet/v6 v6.1.0
	github.com/digitalex70/celeritas v0.0.0-20220222025147-44c03444691c
	github.com/go-chi/chi/v5 v5.0.7
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/alexedwards/scs/v2 v2.4.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
