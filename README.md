# Chipmonk Assignment

## 1. [Usage](#usage)

## 2. [Routes](#routes)

## 3. [Framework](#framework)

## [Usage](#usage)

### Add dependencies

Run the following command to install the necessary dependencies for the REST APIs.

```
go mod download
```

## [Routes](#routes)

1. **`/register`**
   : **`POST`** - use to add new user
2. **`/login`**
   : **`POST`** - used to generate a session token for accessing further routes
3. **`/active-users`**
   : **`GET`** - to fetch the list of active usernames

## [Framework](#framework)

1. Framework used while creating these Project is Gin-Gonic it is HTTP web framework as it is relatively faster than any other frameworks.

2. The Gin framework has it's own Recovery Method, it also has its Default Logger and it also has it Middleware and can also supports any Custom Middleware or Logger.

3. Group Routing is easy in Gin. Authorized routes vs non authorized routes, different API versions can be managed. In addition, the groups can be nested unlimitedly without degrading performance.

4. It can handle JSON, XML or HTML type data for data rendering.
