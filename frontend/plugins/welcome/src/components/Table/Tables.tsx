import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import {Header, Page, pageTheme} from '@backstage/core';
import {Grid,Link} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntDiagnose } from '../../api/models/EntDiagnose';
import { Link as RouterLink } from 'react-router-dom';
const useStyles = makeStyles({
table: {
minWidth: 650,
},
});
export default function ComponentsTable() {
const classes = useStyles();
const api = new DefaultApi();
const [diagnoses, setDiagnoses] = React.useState<EntDiagnose[]>([]);
const [loading, setLoading] = useState(true);
useEffect(() => {
  const getDiagnoses = async () => {
      const res = await api.listDiagnose({ limit: 20, offset: 0 });
      setLoading(false);
      setDiagnoses(res);
      console.log(diagnoses);
      console.log(res);
  };
  getDiagnoses();
}, [loading]);
const deleteDiagnoses = async (id: number) => {
  console.log(id);
  const res = await api.deleteDiagnose({id:id})
  setLoading(true);
  };
const HeaderCustom = {
    minHeight: '50px',
  };
return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`ข้อมูลบันทึกอาการผู้ป่วย`}>
      </Header>
      <Grid item xs={12}></Grid>
<TableContainer component={Paper}>
<Table className={classes.table} aria-label="simple table">
<TableHead>
<TableRow>
<TableCell align="center">ลำดับ</TableCell>
<TableCell align="center">แพทย์</TableCell>
<TableCell align="center">ผู้ป่วย</TableCell>
<TableCell align="center">โรค</TableCell>
<TableCell align="center">อาการ</TableCell>
<TableCell align="center"><Link component={RouterLink} to="/WelcomePage">
           <Button variant="contained" color="primary">
             เพิ่มข้อมูล
           </Button>
         </Link></TableCell>
</TableRow>
</TableHead>
<TableBody>
{diagnoses === undefined
? null
: diagnoses.map((item) => (
<TableRow key={item.id}>
<TableCell align="center">{item.id}</TableCell>
<TableCell align="center">{item.edges?.doctor?.doctorNAME}</TableCell>
<TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
<TableCell align="center">{item.edges?.illness?.illnessName}</TableCell>
<TableCell align="center">{item.edges?.symptoms?.manner}</TableCell>
<TableCell align="center">

<Button
onClick={() => {
if (item.id === undefined){
return;
}

deleteDiagnoses(item.id);
}}
style={{ marginLeft: 10 }}
variant="contained"
color="secondary"
>
Delete
</Button>
</TableCell>
</TableRow>
))}
</TableBody>
</Table>
</TableContainer>
</Page>
);
}