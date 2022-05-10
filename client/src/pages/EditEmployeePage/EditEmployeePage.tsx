import React, {useEffect, useRef, useState} from 'react';
import {
  Box,
  FormControl,
  FormLabel,
  Grid, Heading,
  NumberDecrementStepper, NumberIncrementStepper, NumberInput,
  NumberInputField, NumberInputStepper, Spinner,
} from "@chakra-ui/react";
import Field from '../../../components/Field';
import {useForm} from "react-hook-form";
import Select from "react-select";
import ReactQuill from "react-quill";

import 'react-quill/dist/quill.snow.css';
import CreatableSelect from "react-select/creatable";
import EducationForm from "../../../forms/EducationForm";
import WorkExperienceForm from "../../../forms/WorkExperienceForm";
import employeeStore from "../../../store/employeeStore";
import authStore from "../../../store/authStore";
import {useNavigate} from "react-router-dom";
import RouteTypes from "../../../RouteType";
import {observer} from "mobx-react-lite";

const EditEmployeePage = () => {
  const navigate = useNavigate()
  const { id, type } = authStore
  console.log(JSON.parse(JSON.stringify(authStore)))
  const { employees, isLoading } = employeeStore
  const employee = employeeStore.employees[id || '']

  const [keySkills, setKeySkills] = useState<{skill_name: string}[]>([]);
  const keySkillRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    console.log(id, type)
    if (!id || type !== 'employee') {
      navigate(RouteTypes.Home)
      return
    }

    employeeStore.fetchEmployeeById(id.toString())
  }, []);


  return (
    <Box>
      {
        (isLoading || !employee) && (
          <Spinner size="lg" color="green"/>
        )
      }
    </Box>
  );
};

export default observer(EditEmployeePage);