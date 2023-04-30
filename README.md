# gRPC service for card number validation by Luhn algorithm
Experiment micro gRPC service which implement card number validation using Luhn algorithm.

### Before to start:
- Have to be installed go lang localy, with version not less than 1.18. (service will be dockerized later)

### Clone repo:
```
git clone https://github.com/rubyqorn/luhn-service.git
```
### Create .env file
You can copy example from .env.dist file and create own .env file. This we need because server and client use .env file, but distributed file cant be used
### Change directory and run gRPC server:
```
cd server && go run main.go
```
You will see, that your server have been started and can accept requests from client side:
```
server linstening at: [::]:8080
```
### Change directory and use client:
```
cd client
go run main.go --creditCardNumber=5160215878972826
```
Client script can accept `--creditCardNumber` flag which will be sent to server side and validated by. If flag will not be passed, then client will pass invalid value
### Result:
As result you will see something like this:
```
validation result: true
```
If client gave you `true` this means that card number is valid, otherwise passed card is invalid

### How to generate fake card number?
[Fake card number generator](https://www.creditcardvalidator.org/generator)