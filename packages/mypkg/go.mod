module mypkg

require (
    "service" v0.0.0
)

replace (
    "service" v0.0.0 => "./service"
)

go 1.14
