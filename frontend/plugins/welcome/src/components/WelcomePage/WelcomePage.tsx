import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2'; // alert
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Link,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDoctor } from '../../api/models/EntDoctor'; // import interface User
import { EntIllness } from '../../api/models/EntIllness'; // import interface Video
import { EntSymptoms } from '../../api/models/EntSymptoms'; // import interface Resolution
import { EntPatient } from '../../api/models/EntPatient'; // import interface Playlist
import { WelcomePlugin } from '../../../../../packages/app/src/plugins';


const HeaderCustom = {
  minHeight: '50px',
};
// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));
interface Diagnose {
  illness: number;
  symptoms: number;
  patient: number;
  doctor: number;
}

interface Symptoms {
  manner: string;
}

const Diagnose: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [diagnose, setDiagnose] = React.useState<Partial<Diagnose>>
({});
const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
const [patients, setPatients] = React.useState<EntPatient[]>([]);
const [illnesss, setIllnesss] = React.useState<EntIllness[]>([]);
const [symptomss, setSymptomss] = React.useState<Partial<Symptoms>>({});



  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  }); 


  const getDoctor = async () => {
    const res = await http.listDoctor({ limit: 2, offset: 0 });
    setDoctors(res); 
  };
  const getPatient = async () => {
    const res = await http.listPatient({ limit: 4, offset: 0 });
    setPatients(res);
  };
  const getIllness = async () => {
    const res = await http.listIllness({ limit: 4,  offset: 0 });
    setIllnesss(res);
  };
  
  // Lifecycle Hooks
    useEffect(() => {
    getDoctor();
    getPatient();
    getIllness();
  }, []);

  // set data to object diagnose
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Diagnose;
    const { value } = event.target;
    setDiagnose({ ...diagnose, [name]: value });
    
  };

  const handleChangesymptoms = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Diagnose;
    const { value } = event.target;
    setSymptomss({ ...symptomss, [name]: value });
  };

  // clear input form
  function clear() {
    setDiagnose({});
  }

 // function save data
 function save() {
   
  const apiUrl = 'http://localhost:8080/api/v1/symptomss';
  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(symptomss),
  };
  fetch(apiUrl, requestOptions)
    .then(response => response.json())
    .then(data => {
      diagnose.symptoms = data.id
      console.log(data);
      console.log(diagnose);
      http.createDiagnose({diagnose:diagnose})
      if (data.status === true) {
        clear();
        Toast.fire({
          icon: 'success',
          title: 'บันทึกข้อมูลสำเร็จ',
        });
      } else {
        Toast.fire({
          icon: 'error',
          title: 'บันทึกข้อมูลไม่สำเร็จ',
        });
      }
    });
  }

  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Diagnose`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
        
      </Header>
      <Content>
        <Container maxWidth="sm">
            <Grid item xs={3}></Grid>
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Doctor </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Doctor</InputLabel>
                <Select
                  name="doctor"
                  value={diagnose.doctor } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {doctors.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.doctorNAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            
            <Grid item xs={3}>
              <div className={classes.paper}> Patient </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Patient</InputLabel>
                <Select
                  name="patient"
                  value={diagnose.patient } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}> Illness </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Illness</InputLabel>
                <Select
                  name="illness"
                  value={diagnose.illness } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {illnesss.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.illnessName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}> Symptoms </div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="symptoms"
                  name="manner"
                  type="string"
                  multiline
                  rows={7}
                  style={{ width: 300 }}
                  value={symptomss.manner || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChangesymptoms}
                />
              </form>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
                component={RouterLink}
                to="/Table"
              >
                Save
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};
export default Diagnose;