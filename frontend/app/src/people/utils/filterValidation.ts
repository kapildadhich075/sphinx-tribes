export const bountyHeaderFilter = ({ Paid, Assigned, Open }, bodyPaid, bodyAssignee) => {
  if (Paid) {
    if (Assigned) {
      if (Open) {
        return true;
      } 
        return bodyAssignee || bodyPaid;
      
    } 
      if (Open) {
        return bodyPaid || !bodyAssignee;
      } 
        return bodyPaid;
      
    
  } 
    if (Assigned) {
      if (Open) {
        return !bodyPaid;
      } 
        return !bodyPaid && bodyAssignee;
      
    } 
      if (Open) {
        return !bodyPaid && !bodyAssignee;
      } 
        return true;
      
    
  
};

export const bountyHeaderLanguageFilter = (codingLanguage, filterLanguage) => {
  if (Object.keys(filterLanguage)?.every((key) => !filterLanguage[key])) {
    return true;
  } return codingLanguage?.some(({ value }) => filterLanguage[value]) ?? false;
};
