# pull latest postgres image
FROM postgres:latest


#run init.sql to create the database
ADD init.sql /docker-entrypoint-initdb.d
ADD init_test.sql /docker-entrypoint-initdb.d