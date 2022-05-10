import {EditableEmployeeJobDescription} from "../../types";

export interface EditEmployeeModalProps {
  isOpen: boolean,
  onClose: () => void,
  onSubmit: (data: EditableEmployeeJobDescription, i?: number) => void
  currentIndex?: number,
  employeeJobData?: EditableEmployeeJobDescription | null,
}