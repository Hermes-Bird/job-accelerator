import React, {FC, useState} from 'react';
import {
  Button, Flex, FormControl, FormLabel, Grid, Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay, Textarea
} from "@chakra-ui/react";
import {EditEmployeeModalProps} from "./types";
import {useForm} from "react-hook-form";
import Field from "../Field";
import DatePicker from "react-date-picker";
import {EmployeeJobDescriptionFormData} from "../../types";
import {stringToDate} from "../../helpers/dateHelpers";

const EditEmployeeModal: FC<EditEmployeeModalProps> = ({
  onClose,
  isOpen,
  onSubmit,
  employeeJobData,
  currentIndex,
}) => {
  const {register, handleSubmit } = useForm<EmployeeJobDescriptionFormData>({
    defaultValues: {
      organization: employeeJobData?.organization,
      responsibilities: employeeJobData?.responsibilities
    }
  })
  const [startDate, setStartDate] = useState<Date>(stringToDate(employeeJobData?.start_date) || new Date())
  const [endDate, setEndDate] = useState<Date>(stringToDate(employeeJobData?.end_date) || new Date())

  return (
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay/>
        <ModalContent>
          <ModalHeader>Work description:</ModalHeader>
          <ModalCloseButton/>
          <ModalBody>
            <Grid gap="4">
              <Field
                register={register}
                name="organization"
                label="Organization"
                placeholder="Name of your organization"
              />
              <Flex>
                <FormControl>
                  <FormLabel>Start date</FormLabel>
                  <DatePicker value={startDate} onChange={setStartDate} maxDate={endDate}/>
                </FormControl>
                <FormControl>
                  <FormLabel>End date</FormLabel>
                  <DatePicker value={endDate} onChange={setEndDate} maxDate={new Date()}/>
                </FormControl>
              </Flex>
              <FormControl>
                <FormLabel>Responsibilities</FormLabel>
                <Textarea {...register('responsibilities')} placeholder="Write about your responsibilities"/>
              </FormControl>
            </Grid>
          </ModalBody>
          <ModalFooter>
            <Button onClick={handleSubmit((d) => {
              let initialData = {}
              if (employeeJobData) {
                initialData = employeeJobData
              }
              onSubmit({
                ...initialData,
                ...d,
                start_date: startDate,
                end_date: endDate,
              }, currentIndex)
            })} colorScheme="green">
              Save
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
  );
};

export default EditEmployeeModal;