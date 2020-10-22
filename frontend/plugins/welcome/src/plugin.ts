import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from './components/Login';
import Home from './components/Home';
import Table from './components/Table';

 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login );
    router.registerRoute('/Home', Home );
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Table', Table );
  },
});
