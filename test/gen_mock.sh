#!/bin/bash

mockgen -package=crud -source=app/crud/handler.go -destination=app/crud/handler_mocks_test.go