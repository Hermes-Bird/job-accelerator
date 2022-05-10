import React, {FC} from 'react';
import Field from "../components/Field";
import {useForm} from "react-hook-form";
import {Button, Text} from "@chakra-ui/react";
import * as yup from "yup";
import {yupResolver} from "@hookform/resolvers/yup";
import {LoginDto} from "../types";


const loginEmployeeSchema = yup
  .object()
  .shape({
    email: yup.string().required('email is required').email('invalid email provided'),
    password: yup.string().required('password is required').min(6, 'password should be longer than 5 chars')
  });

const LoginEmployeeForm: FC<{
  onSubmit: (data: LoginDto) => void,
  error?: string
}> = ({ onSubmit, error }) => {
  const { register, handleSubmit, formState: { errors } } = useForm<LoginDto>({
    resolver: yupResolver(loginEmployeeSchema)
  })
  return (
    <>
      <Field
        register={register}
        label="Email"
        name="email"
        error={errors.email?.message}
      />
      <Field
        register={register}
        name="password"
        label="Password"
        type="password"
        error={errors.password?.message}
      />
      {
        error && <Text>{error}</Text>
      }
      <Button onClick={handleSubmit(onSubmit)} colorScheme="linkedin">
        Login
      </Button>
    </>
  );
};

export default LoginEmployeeForm;