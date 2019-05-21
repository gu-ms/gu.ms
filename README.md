# gu.ms
miniature-happiness

A Nature friendly initiative which targets zero-mark-up-language and hence zero-parsing resulting in being CPU (and nature)-friendly. 

Each *gum*, in **gu.ms** is a sticky and short URL that provides some valuable information. For each particular use-case or requirement, 
if there exists a gum that can provide you the required information, it can be obtained by simply calling the respective tiny URL (gum).
These *gums* provide the information as plain text (without any additional mark-up or metadata). This is particulary useful in scripting/development environment, where the code can make use of the gums URL by calling it and directly assigning the return value to a **variable** without any further processing required. 

*gums* could be useful for IOT (and connected embedded) devices where the memory footprint is quite small and cannot have scope to carry libraries and/or binaries to carry out certain actions. If there is a *gum* available for pulling required information, it can help in getting rid of an external library or binary to carry out processing to get the same information.

## Some Use-cases ##

The following is a short list of possible use-cases. Not all might be available yet! 

    Contributors are welcome to add use-cases and/or solutions


### System administrator use-cases ###

    gu.ms/ip               your (public) IP address
    gu.ms/pass             random password generator [default:16 digit]
    gu.ms/pass/an/10       random password generator with 10 digits and *alpha* *numeric* only
    gu.ms/mx/google.com    provides the MX record for google.com domain
    gu.ms/ns/google.com    provides the nameservers for google.com domain
    gu.ms/cert/google.com  provides the certificate validity or where the cert has failed
    gu.ms/man/lsmod        man page of lsmod
    gu.ms/time             provides current time in GMT
    

### Programmer use-cases ###

    gu.ms/client           provides your browser info
    gu.ms/ch/c             provides the ascii value of character 'c'
    gu.ms/bin/93           display binary
    gu.ms/hex/45           display hex
    
### General use-cases ###

    gu.ms/c/usd/gbp        provides cost of 1 USD in GBP
    gu.ms/t/edt            provides current time Eastern Time (DST)
    gu.ms/fc/lon           provides forecast of London
    gu.ms/cv/mi/m/59       converts 59 miles to meters
    gu.ms/fl/a/ba2272      current status/arrival time of flight ba2272
    gu.ms/bmi/100k/165c    BMI calculator

### Use-cases for fun ###
    
    gu.ms/aa/hello         hello converted to ascii art
    gu.ms/morse/hello      hello converted to morse code
    

## Contributors ##
Contributions in ideas, new tiny URL suggestions, code contributions are MOST WELCOME!!


## License
[MIT License](https://raw.githubusercontent.com/gu-ms/gu.ms/master/LICENSE)
