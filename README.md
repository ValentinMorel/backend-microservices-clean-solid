# Back-End Hiring Test Candidate's Response 


This (new) file includes :
- Instructions about how to setup and run the code
- A clear explanation of the approach and design choices


# Setup with docker

## Prerequisites

- docker
- docker-compose

## Run 

In order to run the 2 microservices (gateway and words services), you just need to execute (remove ```-d``` to be directly attached to docker): 

```
docker-compose up -d 
```

to stop it after running in detached mode : 

```docker
docker-compose down 
```

to get logs from all the services : 
 
```
docker-compose logs --tail last.lines
```

or : 

```
docker-compose logs
```

# Setup without docker 

## Prerequisites

- go 1.19

## Run

Change the directory to ```cd path/to/go.mod``` corresponding to the app you want to run, then: 

```
go mod download
```

and run the main.go for the concerned (at least the gateway + 1 service so 2 consoles opened apart): 

```
go run cmd/main.go
```

## Mocks generation for tests

Mock generation needs mockery binary. 

You can install it with : 

```
go install github.com/vektra/mockery/v2@latest
```

and then change directory to the interfaces (that's all we need to build our things) and execute :

```
mockery --name=NameOfTheInterface
```

The result will be placed in ```interfaces/mocks``` folder.


# Approach 

The idea behind this architecture and this approach comes from a full decoupling idea. 
- On the first hand, I believe that microservices should be as independent as they can. One microservice doesn't have to take dependencies from another one.
- In the other hand, the deployed services should be extendable and have reusable pieces of codes to achieve it. (e.g: add a new route to words group)
- Each service have different layers -> (low to top level) database layer, repository layer, usecase layer, delivery layer

The used project layout comes from docker runnable constraint. To keep it simple and light, each Dockerfile is at root of the services then the entire project (including multiple services) doesn't have to be copied everytime we build one service.


# Design choices

The project relies on design patterns to bring various benefits: 

- reducing dependencies between blocks of code to allow an easy and painless replacement
- making the design easier and understandable
- bringing maintainability, testability, scalability and reusability
- bringing bounds to refuse bad designs/anti-patterns by developers extending the software

The software overall follows an onion architecture : the previous layer (lower one) is wrapped in next layer, then uses lower level methods to define higher methods and so on.
Again, the idea is to decouple every aspect of the execution.

## Single Responsibility

Every structure/implementor has one job. One contract with a set of methods. 

## Interface segregation 

The interface segregation allows the software to be highly testable and only expose a small set of client specific methods. 

## Dependency injection 

With this principle, we can make the database easily interchangeable. As long as the database implements the contract (= interface). So a change in code doesn't imply to break it all.

## Dependency inversion

The idea behind the use of this principle is that code should rely on abstractions instead of real implementations.
High level and low level classes should depend on abstractions.

## Open/Closed principle 

We have an example of this principle in the ```WordRepositoryWithFeatures``` structure. 
It wraps the original "implementation" to define a new behavior based on the same abstraction defined earlier. 
So we don't touch the basic methods and add more features by this mechanism.

# What to think about to make it perfect 

- Define a real database instead of a fake array of structures to store data
- Define docker restart policy (e.g: ```restart on failure```)
- Use a context to handle graceful shutdowns 
- Use runtime library to forward informations about the machine usage
- Use a cache between database and requests to decrease database charge and increase response time 
- Comment every function to generate a documentation for further developments

# Self-criticism

- The fake database (array of structs) made the work hard for the design concerning the very low level methods (instead of Query, Scan..., we have a direct access to data) -> it's possible to make it way more cleaner.