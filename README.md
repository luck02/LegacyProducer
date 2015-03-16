# LegacyProducer
Dependancies:
* Follow instructions @ https://github.com/eikonomega/oracle_instant_client_for_ubuntu_64bit To get a working oracle library
* https://registry.hub.docker.com/u/alexeiled/docker-oracle-xe-11g/ 

Copy oci8.pc from <repo>/schema/oci8.pc to /usr/lib/pkgconfig/oci8.pc

https://github.com/mattn/go-oci8.git <-- this should work and you should be able to run the ./DomainPublisher/oracle.go tests.


Import schema:
 * sqlplus64 system/oracle@localhost:49161 @createTables.sql
 * sqlplus64 system/oracle@localhost:49161 @createData.sql
