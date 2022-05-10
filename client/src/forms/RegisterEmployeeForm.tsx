import React, {useState} from 'react';
import {
  Button,
  Radio,
  RadioGroup,
  Stack
} from "@chakra-ui/react";
import {useForm} from "react-hook-form";
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import Field from '../components/Field';

const registerEmployeeSchema = yup
  .object()
  .shape({
    email: yup.string().required('email is required').email('invalid email provided'),
    password: yup.string().min(6, 'password should be longer than 5 chars'),
    first_name: yup.string().required('first name required'),
    last_name: yup.string().required('last name required'),
  })
  .required()


const RegisterCompanyForm = () => {
  const {register, handleSubmit, formState: { errors }} = useForm({
    resolver: yupResolver(registerEmployeeSchema)
  })
  const [sex, setSex] = useState<string>('male');
  console.log(errors)
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
       error={errors.password?.message}
     />
      <Field
        register={register}
        name="first_name"
        label="First name"
        error={errors.first_name?.message}
      />
      <Field
        register={register}
        name="last_name"
        label="Last name"
        error={errors.last_name?.message}
      />
     <RadioGroup value={sex} onChange={setSex}>
        <Stack direction="row">
          <Radio value="male">Male</Radio>
          <Radio value="female">Female</Radio>
        </Stack>
      </RadioGroup>
      <Button
        colorScheme="green"
        onClick={handleSubmit((v) => console.log(v), v => console.log(v))}
        mt={2}
      >
        Register
      </Button>
    </>
  );
};

export default RegisterCompanyForm;