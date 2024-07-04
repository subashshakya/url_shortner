# URL Shortner
### Instructions:
``
git clone https://github.com/subashshakya/url_shortner
``
``
cd url_shortner
``

*For running project*
``go run .``

*For running test*
``go test``

***Project Rundown***
- Storage is implemented as simple map data structure
- HTTP requests are handled by using net/http package from go's standard library
- URL code generation and check for valid URL is provided in utils file
- Implementation of handler functions can be found in handler file
- App runs on localhost port: 3000