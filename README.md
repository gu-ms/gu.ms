# gu.ms
miniature-happiness

A Nature friendly initiative which targets zero-mark-up-language and hence zero-parsing resulting in being CPU (and nature)-friendly. 

Each *gum*, in **gu.ms** is a sticky and short URL that provides some valuable information. For each particular use-case or requirement, 
if there exists a gum that can provide you the required information, it can be obtained by simply calling the respective tiny URL (gum).
These *gums* provide the information as plain text (without any additional mark-up or metadata). This is particulary useful in scripting/development environment, where the code can make use of the gums URL by calling it and directly assigning the return value to a **variable** without any further processing required. 

*gums* could be useful for IOT (and connected embedded) devices where the memory footprint is quite small and cannot have scope to carry libraries and/or binaries to carry out certain actions. If there is a *gum* available for pulling required information, it can help in getting rid of an external library or binary to carry out processing to get the same information.

![calling gu.ms through browser](https://github.com/gu-ms/gu.ms/raw/master/images/browser.png)

![calling gu.ms via curl, lynx, wget on Linux](https://github.com/gu-ms/gu.ms/raw/master/images/tty.png)

![calling gu.ms via curl on Windows](https://github.com/gu-ms/gu.ms/raw/master/images/cmd.png)

## Some Use-cases ##

The following is a short list of possible use-cases. Not all might be available yet! 

    Contributors are welcome to add use-cases and/or gums


### System administrator use-cases ###

URL| Description
---|------------
[gu.ms/ip](http://gu.ms/ip)| your (public) IP address
[gu.ms/pass](http://gu.ms/pass)| random password generator [default:16 digit]
[gu.ms/pass/10/an](http://gu.ms/pass/10/an)| random password generator with 10 digits and *alpha* *numeric* only
[gu.ms/mx/google.com](http://gu.ms/mx/google.com)| the MX record for google.com domain
[gu.ms/ns/google.com](http://gu.ms/ns/google.com)| the nameservers for google.com domain
gu.ms/cert/google.com| provides the certificate validity or where the cert has failed
gu.ms/man/lsmod| man page of lsmod
[gu.ms/time](http://gu.ms/time)| provides current time in GMT
    

##### IP Address #####
    Alias URLs for getting your (public) IP
    
    gu.ms/ip
    gu.ms/whatsmyip
    gu.ms/whatismyip
    gu.ms/myip
    
##### Password Generator ######
    Password generator based on the following classes
    
    1. A: Password must contain ONLY uppercase letters (A...Z)
    2. a: Password must contain ONLY lowercase letters (a...z)
    3. n: Password must contain ONLY digits (or generate a random integer) (0...9)
    4. s: Password must contain ONLY symbols (~.../)
    5. combinations of above: an, An, As, ans... etc
    
    examples:
    
URL| Description
---|------------
[gu.ms/pass/16/a](http://gu.ms/pass/16/a) | generate 16 digit random text using lower case letters only
[gu.ms/pass/10/An](http://gu.ms/pass/10/An) | generate 10 digit random text using upper & lower case letters only
[gu.ms/pass/8](http://gu.ms/pass/8) | generate 8 digit randon password using upper & lower case letters, numbers & symbols
    
##### DNS Records #####
    Several DNS records of a domain can be queried. 
    NOTE: There can be multiple records for each of the query and if present, each record will be on its own line
    
URL|Description
---|------------
[gu.ms/dns/google.com](http://gu.ms/dns/google.com)| fetch IP address of google.com
[gu.ms/mx/example.com](http://gu.ms/mx/example.com)| fetch MX records of domain.com
[gu.ms/ns/example.com](http://gu.ms/ns/example.com)| fetch nameservers
[gu.ms/cname/example.com](http://gu.ms/cname/example.com)| fetch cname records
[gu.ms/srv/example.com/sip/tcp](http://gu.ms/srv/example.com/sip/tcp)| fetch srv records 
[gu.ms/txt/example.com](http://gu.ms/txt/example.com)| fetch txt records
    
    Reverse DNS 
URL| Description
---|------------
[gu.ms/rev/8.8.8.8](http://gu.ms/rev/8.8.8.8)| reverse domain name search (IP address to domain name)     

### Programmer use-cases ###

URL| Description
---|------------
[gu.ms/client](http://gu.ms/client)| provides your browser info
gu.ms/ch/c| provides the ascii value of character 'c'
[gu.ms/bin/93](http://gu.ms/bin/93)| display binary
[gu.ms/hex/45](http://gu.ms/hex/45)| display hex
    
### General use-cases ###

URL| Description
---|------------
gu.ms/c/usd/gbp | provides cost of 1 USD in GBP
gu.ms/t/edt | provides current time Eastern Time (DST)
gu.ms/fc/lon | provides forecast of London
gu.ms/cv/mi/m/59 | converts 59 miles to meters
gu.ms/fl/a/ba2272 | current status/arrival time of flight ba2272
gu.ms/bmi/100k/165c | BMI calculator

### Use-cases for fun ###

URL| Description
---|------------
gu.ms/aa/hello | hello converted to ascii art
gu.ms/morse/hello | hello converted to morse code
    

## Contributors ##
Contributions in ideas, new tiny URL suggestions, code contributions are MOST WELCOME!!


## License
[MIT License](https://raw.githubusercontent.com/gu-ms/gu.ms/master/LICENSE)
