#
# BUILD STAGE
#
FROM maven:3.6.0-jdk-11-slim AS build
COPY .m2 /root/
COPY src /usr/src/app/src  
COPY pom.xml /usr/src/app
RUN mvn -f /usr/src/app/pom.xml clean package

#
# PACKAGE STAGE
#
FROM openjdk:11-jre-slim 
RUN apt-get update && apt-get -y install rsync
COPY --from=build /root/.m2 /root/.m2
COPY --from=build /usr/src/app/target/demo-0.0.1-SNAPSHOT.jar /usr/app/demo-0.0.1-SNAPSHOT.jar  
EXPOSE 8080  
CMD ["java","-jar","/usr/app/demo-0.0.1-SNAPSHOT.jar"]  
