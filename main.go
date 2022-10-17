package main

import (
	"fmt"
	"main/config"
	"main/logging"
)

func writeMessage() {
	fmt.Println("Hello, Platform")
}

func writeMessageLog(logger logging.Logger) {
	logger.Info("Hello, Platform")
}

func writeMessageConfig(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration settings")
		}
	} else {
		logger.Panic("Config section not found")
	}

}
func main() {
	writeMessage()

	// writes out its message using the logging feature
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)
	writeMessageLog(logger)

	// Reading Configuration Settings
	var cfg config.Configuration
	var err error
	cfg, err = config.Load("config.json")
	if err != nil {
		panic(err)
	}

	// DEBUG
	fmt.Println(cfg)

	// # cannot use cfg (type config.Configuration) as type logging.LogLevel in argument to logging.NewDefaultLogger
	// var logger2 logging.Logger = logging.NewDefaultLogger(cfg)
	// writeMessageConfig(logger2, cfg)

}

/*
A custom web application platform (framework) creating features:
logging, session, HTML Templates, authorization.
SportStore App - using a product DB, keeping track of user's product selections,
validating user's input, check out, ...

3-party packages: Gorilla Web Toolkit

1) Logging system
logging.go							//  defines the Logger interface, which specifies methods for logging messages with different levels of severity,
logger_default.go					// the default implementation of the Logger interface Loggers that write messages to standard out
default_create.go					// creates a DefaultLogger with a minimum severity level and log.

2) Configuration System
config.go							// The Configuration interface
config.json							// This configuration file
config_default.go					// Implementing the Configuration Interface
config_default_fallback,go			// To define the methods that accept a default value
config_json.go						// To define the function that will load the data from the configuration file,

3) Managing Services with Dependency Injection
Code depends on an interface can obtain an implementation without needing to select an underlying type or create an instance directly. Код, зависящий от интерфейса, может получить реализацию без необходимости выбирать базовый тип или напрямую создавать экземпляр.

Я предпочитаю использовать внедрение зависимостей (DI), в котором код, зависящий от интерфейса,
может получить реализацию без необходимости выбора базового типа или непосредственного создания экземпляра. я
Я собираюсь начать с определения местоположения службы, которое позже послужит основой для более продвинутых функций.
Во время запуска приложения интерфейсы, определенные приложением, будут добавлены в реестр вместе с
с фабричной функцией, которая создает экземпляры структуры реализации. Так, например, платформа.
Интерфейс logger.Logger будет зарегистрирован с помощью фабричной функции, которая вызывает NewDefaultLogger.
функция. Когда интерфейс добавляется в реестр, он называется службой.
Во время выполнения компоненты приложения, которым нужны функции, описанные службой, переходят в
реестр и запросить интерфейс, который они хотят. Реестр вызывает фабричную функцию и возвращает структуру
который создается, что позволяет компоненту приложения использовать функции интерфейса, не зная или
указав, какая структура реализации будет использоваться или как она создается. Не волнуйтесь, если это не сделает
смысл — это может быть трудной для понимания темой, и становится легче, когда вы видите ее в действии.







**************************************************************
DAO vs Repository patterns?
https://stackoverflow.com/questions/8550124/what-is-the-difference-between-dao-and-repository-patterns
DAO is an abstraction of data persistence. However, a repository is an abstraction of a collection of objects.
DAO is a lower-level concept, closer to the storage systems. However, Repository is a higher-level concept, closer to the Domain objects.

DI: Interfaces->Register, Factory function FF => Sevices+Lifecycles when to invoke FF to create new structure

The Service Lifecycles
Lifecycle 	Description
Transient 	For this lifecycle, the factory function is invoked for every service request.
Singleton 	For this lifecycle, the factory function is invoked once, and every request receives the same struct instance.
Scoped 		For every request within that scope receives the same struct instance.

lifecycles.go
context.go					// to work with contexts (Ch. 30)


DAO (OBJECTS to access data)					Repository (Domain Model terms, not DB terms) - contains methods
/to get data simpler then queries/
DAO is an abstraction of data persistence.			Repository is an abstraction of a collection of objects.
considered closer to the database, often table-centric		considered closer to the Domain, dealing only in Aggregate Roots.
Collection<Permission> findPForUser(uId string)			It's a repository of a specific type of objects to handle
User findUser(uId string)					AppleR.findAll(criteria), AppleR.save(juicyApple) - ONE Type of objects only
Collection<User> findUsersForP(premission P)

Repository could be implemented using DAO's, but you wouldn't do the opposite.
Also, a Repository is generally a narrower interface. It should be simply a collection of objects, with a Get(id), Find(ISpecification), Add(Entity).
A method like Update is appropriate on a DAO, but not a Repository - when using a Repository, changes to entities would usually be tracked by separate UnitOfWork.

Note that both patterns really mean the same (they store data and they abstract the access to it and they are both expressed closer to the domain model and
hardly contain any DB reference), but the way they are used can be slightly different, DAO being a bit more flexible/generic, while Repository is a bit more
specific and restrictive to a type only.


*/
