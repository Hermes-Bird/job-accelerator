import React from 'react';
import {
  Box,
  Button,
  Container,
  Divider,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Image,
  Input,
  Link
} from "@chakra-ui/react";
import employeeBg from '../../static/ibiza.jpg'
import searchEmployees from '../../static/search-employee.png'

const Login = () => {
  return (
    <Box>
      <Box
        pos="absolute"
        top="0"
        bottom="0"
        right="0"
        left="0"
        backgroundImage={employeeBg}
        bgSize="cover"
        filter="blur(0px)"
        zIndex={-1}
        h="100%"
      />
      <Container maxW={"container.sm"}>
        <Box
          zIndex={1}
          bg="white"
          borderRadius={5}
          border="3px solid"
          borderColor="green.200"
          p={5}
          boxShadow="1px 3px 10px black"
        >
          <Heading textAlign="center" fontStyle={"italic"} fontWeight={400} mb={4}>Login as employee</Heading>
          <Flex flexDir="column" gap={4}>
            <FormControl>
              <FormLabel>Email</FormLabel>
              <Input type="email"/>
            </FormControl>
            <FormControl>
              <FormLabel>Password</FormLabel>
              <Input type="password"/>
            </FormControl>
            <Button
              colorScheme="green"
              mt={2}
            >
              Login
            </Button>
            <Divider mt={3}/>
            <Link textAlign="center" textColor="teal.300">
              I already have an account
            </Link>
          </Flex>
        </Box>
        <Box bg="white" border="3px solid" borderColor="teal.300" borderRadius={5}>
          <Heading fontWeight={400} size={"sm"}>I search for employees</Heading>
          <Image src={searchEmployees} w="100" h="100" />
        </Box>
      </Container>
  </Box>
  );
};

export default Login;