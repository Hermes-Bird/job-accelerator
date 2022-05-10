export const experienceCodeToText = (expCode: string): string => {
  if (expCode.startsWith('0')) {
    return 'without experience'
  }
  if (expCode.startsWith('1')) {
    return '1 year'
  }
  if (expCode.endsWith('+')) {
    return 'more than 5 years'
  }

  return `${expCode[1]}  years`
}