module example.com/mymodule

go 1.14
toolchain go1.21

require (
    example.com/othermodule v1.2.3
    example.com/thismodule v1.6.3
    example.com/thatmodule v1.1.3
)

require (
    example.com/anothermodule v1.7.3 // indirect
)

replace example.com/thatmodule => ../thatmodule
replace example.com/amodule v1.2.3 => example.com/amodule v1.2.4
exclude example.com/thismodule v1.3.0

retract v1.1.0 // broken
retract [v1.1.2,v1.1.5] // bug