# Basic GO microservice with Chi

## Description

This is a basic microservice written in GO using Chi framework. It uses Redis as a database. It has a simple CRUD for orders.

## apis

- Home api
  `/`

- Order apis

  `/orders` - GET - get all orders

  `/orders` - POST - create new order

  `/orders/{id}` - GET - get order by id

  `/orders/{id}` - PUT - update order by id

  `/orders/{id}` - DELETE - delete order by id
